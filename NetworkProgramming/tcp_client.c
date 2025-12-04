#include <string.h>
#include <sys/types.h>
#include <unistd.h> // close
#include <arpa/inet.h> // htons, inet_addr()
#include <stdlib.h>
#include <stdio.h>
#include <strings.h> //bezero()
#include <sys/socket.h>
#define PORT 8080
#define BUFFER_SIZE 1024
#define SoAd struct sockaddr


void chat(int sockfd){
    char buffer[BUFFER_SIZE] = {0};
    printf("If you want to end communicate, enter: exit\n");
    int n = 0;
    while (1) {
        // bzero(buffer, sizeof(buffer));
        memset(buffer, 0, BUFFER_SIZE);
        printf("Enter the Message : ");
        n = 0;
        while((buffer[n++] = getchar()) != '\n' && n< BUFFER_SIZE-1);
        buffer[n] = '\0';

        // don't read/write the full BUFFER_SIZE, this wastes bandwith.
        // write(sockfd, buffer, sizeof(buffer));  

        // size_t len = strlen(buffer);
        // write(sockfd, buffer, len);
        // use send instand of write
        if(send(sockfd, buffer, n, 0)<0){
            perror("send faild.");
            break;
        }

        // bzero(buffer, sizeof(buffer));
        memset(buffer, 0, BUFFER_SIZE);

        // read(sockfd, buffer, BUFFER_SIZE-1);

        // handle return values
        /*ssize_t bytes_read = read(sockfd, buffer, BUFFER_SIZE - 1);
        if (bytes_read <= 0) {
            // Handle error or connection closed
            printf("read failed or connection closed.\n");
            break;
        }
        buffer[bytes_read] = '\0';  // Null-terminate
        */

        // use recv instand of read
        ssize_t bytes_receive = recv(sockfd, buffer, BUFFER_SIZE-1, 0);
        if(bytes_receive <= 0) {
            if(bytes_receive < 0) perror("receive failed.");
            else printf("Server disconnected.\n");
            break;
        }
        buffer[bytes_receive] = '\0';


        printf("Server : %s",buffer);
        if(strcmp(buffer, "exit\n") ==0 ){
            printf("Client Exit...\n");
            break;
        }
    }
}


int main(){
    int sockfd;
    struct sockaddr_in servaddr;


    sockfd  = socket(AF_INET, SOCK_STREAM, 0);
    if(sockfd == -1){
        perror("socket creation failed.");
        exit(0);
    }

    // bzero(&servaddr, sizeof(servaddr));
    memset(&servaddr, 0, sizeof(servaddr));

    // assign IP, PORT
    servaddr.sin_family = AF_INET;
    servaddr.sin_port =  htons(PORT);// arpa/inet.h
    
    // Convert IPv4 numbers-and-dots notation into binary form (in network  byte order)
    servaddr.sin_addr.s_addr = inet_addr("127.0.0.1");// arpa/inet.h

    if(connect(sockfd, (SoAd*)&servaddr, sizeof(servaddr)) != 0){
        perror("connection with server failed.");
        close(sockfd);
        exit(0);
    }

    printf("connection succeed.\n");

    // chat
    chat(sockfd);

    close(sockfd);
    return 0;
}