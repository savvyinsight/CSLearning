#include <pthread.h>
#include <arpa/inet.h>
#include <netinet/in.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <strings.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <time.h>
#include <unistd.h>

#define BUFFER_SIZE 1024
#define PORT 8080
#define SoAd struct sockaddr

void chat(int connfd){
    char buffer[BUFFER_SIZE];
    int n = 0;
    while(1){
        memset(buffer, 0, BUFFER_SIZE);

        ssize_t byte_received = recv(connfd, buffer, BUFFER_SIZE-1, 0);
        if(byte_received <= 0) {
            if(byte_received < 0) perror("receive failed.");
            else printf("Client disconnected.\n");
            break;  // Exit loop!
        }
        buffer[byte_received] = '\0';

        // printf buffer which contains the client contents
        printf("From Client : %s ",buffer);

        // bzero(buffer, BUFFER_SIZE);
        memset(buffer, 0, BUFFER_SIZE);

        n = 0;
        printf("To Client : ");
        // copy server message in the buffer
        while((buffer[n++] = getchar()) != '\n' && n<BUFFER_SIZE-1);
        buffer[n] = '\0'; // Null-termination

        // use send instand of write
        if(send(connfd, buffer, n, 0)<0){
            perror("send failed.");
            break;
        }

        if(strcmp(buffer, "exit\n") == 0){
            printf("Server Exit.\n");
            break;  
        }
    }
    // server close connfd
    close(connfd);
}

// Thread handler function
void *client_handler(void *arg){
    int connfd = *((int*)arg);
    free(arg);  // Free the allocated memory
    
    // Call chat function
    chat(connfd);
    
    pthread_exit(NULL);
}

int main(){
    int sockfd,connfd;
    struct sockaddr_in servaddr, cliaddr;

    sockfd = socket(AF_INET, SOCK_STREAM, 0);
    if(sockfd == -1){
        perror("socket creation failed.");
        exit(0);
    }

    memset(&servaddr, 0, sizeof(servaddr));

    // assign IP, PORT
    servaddr.sin_family = AF_INET;
    servaddr.sin_addr.s_addr = INADDR_ANY;
    servaddr.sin_port = htons(PORT);

    if(bind(sockfd, (SoAd*)&servaddr, sizeof(servaddr))
        != 0){
        perror("socket bind failed.");
        close(sockfd);
        exit(0);        
    }

    if(listen(sockfd, 5) != 0){
        perror("listen failed.");
        close(sockfd);
        exit(0);
    }

    printf("Server listening .... \n");

    while(1){
        socklen_t len = sizeof(cliaddr);
        connfd = accept(sockfd, (SoAd*)&cliaddr, &len);
        if(connfd < 0){
            perror("server accept failed");
            close(connfd);
            // exit(0);
            continue; // Don't exit,just continue;
        }

        printf("Received connection from %s:%d\n", 
               inet_ntoa(cliaddr.sin_addr), ntohs(cliaddr.sin_port));

        // Create a thread for this client
        pthread_t thread_id;
        int *client_fd = malloc(sizeof(int));
        *client_fd = connfd;  // Pass the socket descriptor

        if(pthread_create(&thread_id, NULL, client_handler, client_fd) != 0){
            perror("Failed to create thread");
            free(client_fd);
            close(connfd);
            continue;
        }

        // Detach thread so resources are automatically reclaimed
        pthread_detach(thread_id);
    }

    close(sockfd);

    return 0;
}