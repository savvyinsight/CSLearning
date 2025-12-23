#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <errno.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <sys/epoll.h>



#define PORT 8080
#define MAX_EVENTS 1024
#define BUFFER_SIZE 4096
#define MAX_CLIENTS 10000


static int set_nonblocking(int fd) {
    int flags = fcntl(fd, F_GETFL, 0);
    if (flags == -1) {
        perror("fcntl F_GETFL");
        return -1;
    }
    
    if (fcntl(fd, F_SETFL, flags | O_NONBLOCK) == -1) {
        perror("fcntl F_SETFL");
        return -1;
    }
    
    return 0;
}

// Client connection structure
typedef struct {
    int fd;
    struct sockaddr_in addr;
    char buffer[BUFFER_SIZE];
    size_t buf_len;
    int need_write;  // Flag to indicate if we need to write to this client
} client_t;


static void init_client(client_t *client, int fd, struct sockaddr_in *addr) {
    client->fd = fd;
    client->addr = *addr;
    memset(client->buffer, 0, BUFFER_SIZE);
    client->buf_len = 0;
    client->need_write = 0;
}

int main() {
    int server_fd, epoll_fd;
    struct sockaddr_in server_addr;
    struct epoll_event ev, events[MAX_EVENTS];
    // client_t clients[MAX_CLIENTS];
    int nfds;
    
    // Dynamic memory allocation
    client_t *clients = malloc(MAX_CLIENTS * sizeof(client_t));
    if (!clients) {
        perror("malloc clients failed");
        exit(EXIT_FAILURE);
    }

    for (int i = 0; i < MAX_CLIENTS; i++) {
        clients[i].fd = -1;  // Mark as unused
    }
    
    if ((server_fd = socket(AF_INET, SOCK_STREAM | SOCK_NONBLOCK, 0)) == -1) {
        perror("socket failed");
        exit(EXIT_FAILURE);
    }
    
    int opt = 1;
    if (setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt))) {
        perror("setsockopt failed");
        exit(EXIT_FAILURE);
    }
    
    server_addr.sin_family = AF_INET;
    server_addr.sin_addr.s_addr = INADDR_ANY;
    server_addr.sin_port = htons(PORT);
    
    if (bind(server_fd, (struct sockaddr*)&server_addr, sizeof(server_addr)) == -1) {
        perror("bind failed");
        close(server_fd);
        exit(EXIT_FAILURE);
    }

    if (listen(server_fd, SOMAXCONN) == -1) {
        perror("listen failed");
        close(server_fd);
        exit(EXIT_FAILURE);
    }
    
    epoll_fd = epoll_create1(0);
    if (epoll_fd == -1) {
        perror("epoll_create1 failed");
        close(server_fd);
        exit(EXIT_FAILURE);
    }
    
    // Add server socket to epoll
    ev.events = EPOLLIN | EPOLLET;  // Edge-triggered mode
    ev.data.fd = server_fd;
    if (epoll_ctl(epoll_fd, EPOLL_CTL_ADD, server_fd, &ev) == -1) {
        perror("epoll_ctl: server_fd");
        close(epoll_fd);
        close(server_fd);
        exit(EXIT_FAILURE);
    }
    
    printf("Epoll echo server listening on port %d\n", PORT);
    printf("Maximum connections: %d\n", MAX_CLIENTS);
    printf("Using edge-triggered mode\n");
    
    while (1) {
        nfds = epoll_wait(epoll_fd, events, MAX_EVENTS, -1);
        if (nfds == -1) {
            perror("epoll_wait");
            break;
        }
        
        // Process all ready events
        for (int i = 0; i < nfds; i++) {
            // Handle new connections
            if (events[i].data.fd == server_fd) {
                // Accept all incoming connections
                while (1) {
                    struct sockaddr_in client_addr;
                    socklen_t client_len = sizeof(client_addr);
                    int client_fd = accept(server_fd, 
                                           (struct sockaddr*)&client_addr,
                                           &client_len);
                    
                    set_nonblocking(client_fd);
                    
                    if (client_fd == -1) {
                        if (errno == EAGAIN || errno == EWOULDBLOCK) {
                            // No more connections pending
                            break;
                        } else {
                            perror("accept4 failed");
                            break;
                        }
                    }
                    
                    // Find free slot for client
                    int client_index = -1;
                    for (int j = 0; j < MAX_CLIENTS; j++) {
                        if (clients[j].fd == -1) {
                            client_index = j;
                            init_client(&clients[j], client_fd, &client_addr);
                            break;
                        }
                    }
                    
                    if (client_index == -1) {
                        printf("Too many connections, rejecting new client\n");
                        close(client_fd);
                        continue;
                    }
                    
                    printf("New connection from %s:%d (slot: %d, fd: %d)\n",
                           inet_ntoa(client_addr.sin_addr),
                           ntohs(client_addr.sin_port),
                           client_index,
                           client_fd);
                    
                    // Add client to epoll with edge-triggered mode
                    ev.events = EPOLLIN | EPOLLET | EPOLLRDHUP | EPOLLERR;
                    ev.data.ptr = &clients[client_index];  // Store pointer to client struct
                    if (epoll_ctl(epoll_fd, EPOLL_CTL_ADD, client_fd, &ev) == -1) {
                        perror("epoll_ctl: client_fd");
                        close(client_fd);
                        clients[client_index].fd = -1;
                        continue;
                    }
                }
            }
            // Handle client events
            else {
                client_t *client = (client_t*)events[i].data.ptr;
                
                // Check for connection closure or error
                if (events[i].events & (EPOLLRDHUP | EPOLLHUP | EPOLLERR)) {
                    printf("Client %s:%d disconnected\n",
                           inet_ntoa(client->addr.sin_addr),
                           ntohs(client->addr.sin_port));
                    
                    // Remove from epoll
                    epoll_ctl(epoll_fd, EPOLL_CTL_DEL, client->fd, NULL);
                    
                    // Close socket and mark slot as free
                    close(client->fd);
                    client->fd = -1;
                    continue;
                }
                
                // Handle read events (edge-triggered: must read ALL available data)
                if (events[i].events & EPOLLIN) {
                    // Read all available data
                    while (1) {
                        ssize_t bytes_read = recv(client->fd,
                                                client->buffer + client->buf_len,
                                                BUFFER_SIZE - client->buf_len - 1,
                                                0);
                        
                        if (bytes_read > 0) {
                            client->buf_len += bytes_read;
                            client->buffer[client->buf_len] = '\0';
                            
                            printf("Received %zd bytes from %s:%d (total buffered: %zu)\n",
                                   bytes_read,
                                   inet_ntoa(client->addr.sin_addr),
                                   ntohs(client->addr.sin_port),
                                   client->buf_len);
                            
                            // Mark that we need to write back to this client
                            client->need_write = 1;
                            
                            // If buffer is getting full, echo immediately
                            if (client->buf_len >= BUFFER_SIZE / 2) {
                                // Update epoll to also watch for write readiness
                                ev.events = EPOLLIN | EPOLLET | EPOLLRDHUP | EPOLLERR | EPOLLOUT;
                                ev.data.ptr = client;
                                epoll_ctl(epoll_fd, EPOLL_CTL_MOD, client->fd, &ev);
                            }
                        } 
                        else if (bytes_read == 0) {
                            // Connection closed by client
                            printf("Client %s:%d closed connection\n",
                                   inet_ntoa(client->addr.sin_addr),
                                   ntohs(client->addr.sin_port));
                            
                            epoll_ctl(epoll_fd, EPOLL_CTL_DEL, client->fd, NULL);
                            close(client->fd);
                            client->fd = -1;
                            break;
                        }
                        else {
                            if (errno == EAGAIN || errno == EWOULDBLOCK) {
                                // No more data to read (edge-triggered behavior)
                                break;
                            } else {
                                perror("recv failed");
                                
                                epoll_ctl(epoll_fd, EPOLL_CTL_DEL, client->fd, NULL);
                                close(client->fd);
                                client->fd = -1;
                                break;
                            }
                        }
                    }
                }
                
                // Handle write events
                if (events[i].events & EPOLLOUT) {
                    if (client->need_write && client->buf_len > 0) {
                        ssize_t bytes_sent = send(client->fd,
                                                 client->buffer,
                                                 client->buf_len,
                                                 0);
                        
                        if (bytes_sent > 0) {
                            printf("Echoed %zd bytes to %s:%d\n",
                                   bytes_sent,
                                   inet_ntoa(client->addr.sin_addr),
                                   ntohs(client->addr.sin_port));
                            
                            // Remove sent data from buffer
                            if (bytes_sent < client->buf_len) {
                                // Partial write, move remaining data to front
                                memmove(client->buffer,
                                        client->buffer + bytes_sent,
                                        client->buf_len - bytes_sent);
                                client->buf_len -= bytes_sent;
                                client->buffer[client->buf_len] = '\0';
                            } else {
                                // All data sent
                                client->buf_len = 0;
                                client->need_write = 0;
                                
                                // Remove EPOLLOUT from events if no more data to send
                                ev.events = EPOLLIN | EPOLLET | EPOLLRDHUP | EPOLLERR;
                                ev.data.ptr = client;
                                epoll_ctl(epoll_fd, EPOLL_CTL_MOD, client->fd, &ev);
                            }
                        }
                        else if (bytes_sent == -1) {
                            if (errno == EAGAIN || errno == EWOULDBLOCK) {
                                // Send buffer full, will retry on next EPOLLOUT
                                printf("Send buffer full for %s:%d\n",
                                       inet_ntoa(client->addr.sin_addr),
                                       ntohs(client->addr.sin_port));
                            } else {
                                perror("send failed");
                                
                                epoll_ctl(epoll_fd, EPOLL_CTL_DEL, client->fd, NULL);
                                close(client->fd);
                                client->fd = -1;
                            }
                        }
                    } else {
                        // No data to send, remove EPOLLOUT
                        ev.events = EPOLLIN | EPOLLET | EPOLLRDHUP | EPOLLERR;
                        ev.data.ptr = client;
                        epoll_ctl(epoll_fd, EPOLL_CTL_MOD, client->fd, &ev);
                    }
                }
            }
        }
        
        // Optional: Add statistics or monitoring here
        static unsigned long long total_events = 0;
        total_events += nfds;
        if (total_events % 1000 == 0) {
            // Count active connections
            int active_clients = 0;
            for (int i = 0; i < MAX_CLIENTS; i++) {
                if (clients[i].fd != -1) active_clients++;
            }
            printf("Processed %llu total events, active clients: %d\n",
                   total_events, active_clients);
        }
    }
    
    // Cleanup
    for (int i = 0; i < MAX_CLIENTS; i++) {
        if (clients[i].fd != -1) {
            close(clients[i].fd);
        }
    }
    close(server_fd);
    close(epoll_fd);
    
    return 0;
}