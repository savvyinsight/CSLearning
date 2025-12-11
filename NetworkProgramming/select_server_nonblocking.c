#include <fcntl.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <errno.h>
#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <sys/select.h>
#include <sys/socket.h>
#define PORT 8080
#define BUFFER_SIZE 1024
#define MAX_CLIENTS 24

int set_nonblocking(int sockfd);

int main(){
    int server_fd,new_socket;
    char buffer[BUFFER_SIZE] = {0};
    struct sockaddr_in serveaddr;
    int opt = 1;

    // Master file descriptor sets
    fd_set readfds,masterfds;
    int max_fd;

    // Client sockets array
    int client_sockets[MAX_CLIENTS] = {0};

    if((server_fd = socket(AF_INET, SOCK_STREAM, 0))==0){
        perror("socket creation failed");
        exit(EXIT_FAILURE);
    }

    // Set Non-Blocking mode
    if(set_nonblocking(server_fd) < 0){
        close(server_fd);
        exit(1);
    }

    // Set socket options
    if (setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt))) {
        perror("setsockopt");
        exit(EXIT_FAILURE);
    }

    serveaddr.sin_family = AF_INET;
    serveaddr.sin_addr.s_addr = INADDR_ANY;
    serveaddr.sin_port = htons(PORT);

    if(bind(server_fd, (struct sockaddr*)&serveaddr, sizeof(serveaddr))<0){
        perror("Bind failed");
        exit(EXIT_FAILURE);
    }

    if(listen(server_fd, 5)<0){
        perror("listen");
        exit(EXIT_FAILURE);
    }

    printf("Slect server listening on port:%d\n",PORT);

    // Initialize master set
    FD_ZERO(&masterfds);
    FD_SET(server_fd, &masterfds);
    max_fd = server_fd;

    while (1) {
        // Copy master set to working set (select modifies the set)
        readfds = masterfds;
        
        // Wait for activity on any socket
        printf("Waiting on select()...\n");
        int activity = select(max_fd + 1, &readfds, NULL, NULL, NULL);
        
        if ((activity < 0) && (errno != EINTR)) {// EINTR means "interrupted by signal" - not a real error, just retry
            perror("select error");
        }

        socklen_t server_len = sizeof(serveaddr);


        // Check if server socket has new connection
        if (FD_ISSET(server_fd, &readfds)) {
            if ((new_socket = accept(server_fd, 
                                     (struct sockaddr *)&serveaddr, 
                                     (socklen_t*)&server_len)) < 0) {
                perror("accept");
                exit(EXIT_FAILURE);
            }

            // Set Non-Blocking
            if(set_nonblocking(new_socket)<0){
                close(new_socket);
                continue;
            }
            
            printf("New connection, socket fd: %d, IP: %s, Port: %d\n",
                   new_socket, inet_ntoa(serveaddr.sin_addr), ntohs(serveaddr.sin_port));
            
            // Add new socket to client array
            for (int i = 0; i < MAX_CLIENTS; i++) {
                if (client_sockets[i] == 0) {
                    client_sockets[i] = new_socket;
                    printf("Adding to list of sockets at index %d\n", i);
                    break;
                }
            }
            
            // Add new socket to master set
            FD_SET(new_socket, &masterfds);
            
            // Update max_fd if needed
            if (new_socket > max_fd) {
                max_fd = new_socket;
            }
        }

        // Check all client sockets for data
        for (int i = 0; i < MAX_CLIENTS; i++) {
            int sd = client_sockets[i];
            
            if (sd > 0 && FD_ISSET(sd, &readfds)) {
                // Read incoming data
                int valread = read(sd, buffer, BUFFER_SIZE);
                
                if (valread == 0) {
                    // Client disconnected
                    getpeername(sd, (struct sockaddr*)&serveaddr, 
                                (socklen_t*)&serveaddr);
                    printf("Host disconnected, IP: %s, Port: %d\n",
                           inet_ntoa(serveaddr.sin_addr), ntohs(serveaddr.sin_port));
                    
                    // Close socket and clean up
                    close(sd);
                    FD_CLR(sd, &masterfds);
                    client_sockets[i] = 0;
                }else if(valread < 0){
                    // Check if this is a "would block" error (normal for non-blocking)
                    if (errno == EAGAIN || errno == EWOULDBLOCK) {
                        // No data available right now - this is normal
                        // Just continue to next socket
                        continue;
                    } else {
                        // Real error
                        perror("read error");
                        close(sd);
                        FD_CLR(sd, &masterfds);
                        client_sockets[i] = 0;
                    }
                }else {
                    // Echo back the data
                    buffer[valread] = '\0';
                    printf("Received from client %d: %s\n", sd, buffer);
                    
                    // Send echo back
                    send(sd, buffer, strlen(buffer), 0);
                    printf("Echo sent to client %d\n", sd);
                }
            }
        }
    }

    return 0;
}

int set_nonblocking(int sockfd){
    int flags = fcntl(sockfd,F_GETFL,0);
    if(flags < 0){
        perror("fcntl F_GETFL");
        return -1;
    }

    flags |= O_NONBLOCK;
    if(fcntl(sockfd, F_SETFL,flags) < 0){
        perror("fcntl F_SETFL");
        return -1;
    }
    return 0;
}