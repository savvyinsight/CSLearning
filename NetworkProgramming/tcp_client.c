#include <string.h>
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
        bzero(buffer, sizeof(buffer));
        printf("Enter the Message : ");
        n = 0;
        while((buffer[n++] = getchar()) != '\n');

        write(sockfd, buffer, sizeof(buffer));
        bzero(buffer, sizeof(buffer));
        read(sockfd, buffer, sizeof(buffer));
        printf("Server : %s",buffer);
        if((strncmp(buffer, "exit", 4)) ==0 ){
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

    bzero(&servaddr, sizeof(servaddr));

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