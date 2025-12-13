#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <poll.h>
#include <errno.h>

#define PORT 8080
#define MAX_CLIENTS 100
#define BUFFER_SIZE 1024
#define TIMEOUT 1000  // 1 second timeout

// Helper function to set socket to non-blocking mode
int set_nonblocking(int fd) {
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

int main() {
    int server_fd, client_fd;
    struct sockaddr_in server_addr, client_addr;
    socklen_t client_len = sizeof(client_addr);
    struct pollfd fds[MAX_CLIENTS + 1];  // +1 for server socket
    int nfds = 1;  // Start with only server socket
    char buffer[BUFFER_SIZE];
    
    // Create server socket
    if ((server_fd = socket(AF_INET, SOCK_STREAM, 0)) == -1) {
        perror("socket failed");
        exit(EXIT_FAILURE);
    }
    
    // Set server socket to non-blocking
    if (set_nonblocking(server_fd) == -1) {
        close(server_fd);
        exit(EXIT_FAILURE);
    }
    
    // Set socket options
    int opt = 1;
    if (setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt))) {
        perror("setsockopt failed");
        exit(EXIT_FAILURE);
    }
    
    // Configure server address
    server_addr.sin_family = AF_INET;
    server_addr.sin_addr.s_addr = INADDR_ANY;
    server_addr.sin_port = htons(PORT);
    
    // Bind socket
    if (bind(server_fd, (struct sockaddr*)&server_addr, sizeof(server_addr)) == -1) {
        perror("bind failed");
        close(server_fd);
        exit(EXIT_FAILURE);
    }
    
    // Listen for connections
    if (listen(server_fd, 10) == -1) {
        perror("listen failed");
        close(server_fd);
        exit(EXIT_FAILURE);
    }
    
    printf("Non-blocking echo server listening on port %d\n", PORT);
    printf("Timeout set to %d ms\n", TIMEOUT);
    
    // Initialize pollfd array
    fds[0].fd = server_fd;
    fds[0].events = POLLIN;  // We're interested in read events
    
    // Initialize all other fds to -1 (unused)
    for (int i = 1; i <= MAX_CLIENTS; i++) {
        fds[i].fd = -1;
        fds[i].events = POLLIN;
    }
    
    // Main server loop
    while (1) {
        // Wait for events with timeout
        int ret = poll(fds, nfds, TIMEOUT);
        
        if (ret == -1) {
            perror("poll failed");
            break;
        }
        
        if (ret == 0) {
            // Timeout occurred - no events, can do other work here
            // printf("Poll timeout - no events\n");
            continue;
        }
        
        // Check server socket for new connections
        if (fds[0].revents & POLLIN) {
            // Accept all pending connections
            while (1) {
                client_fd = accept(server_fd, (struct sockaddr*)&client_addr, &client_len);
                
                if (client_fd == -1) {
                    if (errno == EWOULDBLOCK || errno == EAGAIN) {
                        // No more connections pending
                        break;
                    } else {
                        perror("accept failed");
                        break;
                    }
                }
                
                // Set client socket to non-blocking
                if (set_nonblocking(client_fd) == -1) {
                    close(client_fd);
                    continue;
                }
                
                printf("New connection from %s:%d (fd: %d)\n", 
                       inet_ntoa(client_addr.sin_addr), 
                       ntohs(client_addr.sin_port),
                       client_fd);
                
                // Find empty slot in pollfd array
                int slot = -1;
                for (int i = 1; i <= MAX_CLIENTS; i++) {
                    if (fds[i].fd == -1) {
                        slot = i;
                        fds[i].fd = client_fd;
                        fds[i].events = POLLIN;  // Monitor for read events
                        break;
                    }
                }
                
                if (slot == -1) {
                    printf("Too many connections, closing new connection\n");
                    close(client_fd);
                } else {
                    printf("Client added to slot %d\n", slot);
                    if (slot + 1 > nfds) {
                        nfds = slot + 1;
                    }
                }
            }
        }
        
        // Check all client sockets for data
        for (int i = 1; i < nfds; i++) {
            if (fds[i].fd == -1) continue;
            
            int client_socket = fds[i].fd;
            
            // Check for read events
            if (fds[i].revents & POLLIN) {
                // Read all available data from client
                while (1) {
                    memset(buffer, 0, BUFFER_SIZE);
                    ssize_t bytes_read = recv(client_socket, buffer, BUFFER_SIZE - 1, 0);
                    
                    if (bytes_read > 0) {
                        // Successfully read data
                        buffer[bytes_read] = '\0';
                        printf("Received %zd bytes from client %d: %s", 
                               bytes_read, i, buffer);
                        
                        // Echo the data back
                        ssize_t bytes_sent = send(client_socket, buffer, bytes_read, 0);
                        if (bytes_sent == -1) {
                            if (errno == EWOULDBLOCK || errno == EAGAIN) {
                                printf("Send buffer full for client %d, will retry\n", i);
                                // Could store data for later sending
                            } else {
                                perror("send failed");
                                close(client_socket);
                                fds[i].fd = -1;
                                break;
                            }
                        } else {
                            printf("Echoed %zd bytes back to client %d\n", bytes_sent, i);
                        }
                    } 
                    else if (bytes_read == 0) {
                        // Connection closed by client
                        printf("Client %d disconnected\n", i);
                        close(client_socket);
                        fds[i].fd = -1;
                        break;
                    } 
                    else {
                        // Error or no more data
                        if (errno == EWOULDBLOCK || errno == EAGAIN) {
                            // No more data to read
                            break;
                        } else {
                            perror("recv failed");
                            close(client_socket);
                            fds[i].fd = -1;
                            break;
                        }
                    }
                }
            }
            
            // Handle write events if we need to send buffered data
            if (fds[i].revents & POLLOUT) {
                // If we were buffering data for sending, send it here
                // For echo server, we send immediately, so not much to do here
                printf("Client %d socket ready for writing\n", i);
            }
            
            // Handle error events
            if (fds[i].revents & (POLLERR | POLLHUP | POLLNVAL)) {
                printf("Client %d error/disconnect (events: 0x%x)\n", i, fds[i].revents);
                close(client_socket);
                fds[i].fd = -1;
            }
        }
        
        // Optional: Do other work here when poll times out or between events
        // This is where non-blocking I/O shines - you can process other tasks
        // while waiting for I/O events
        
        // Example: periodic tasks
        static int counter = 0;
        if (ret == 0) {  // Poll timeout
            counter++;
            if (counter % 10 == 0) {
                printf("Server idle, doing background work...\n");
            }
        }
    }
    
    // Cleanup
    for (int i = 0; i < nfds; i++) {
        if (fds[i].fd != -1) {
            close(fds[i].fd);
        }
    }
    
    return 0;
}