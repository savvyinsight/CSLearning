#include <stdio.h>
#include <stdlib.h>
#include <netinet/in.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

#define BUFFER_SIZE 1024
#define PORT 8080
int main(){
    int sockfd;
    struct sockaddr_in servaddr;
    char buffer[BUFFER_SIZE];
    const char* hello = "Hello From Client.";

    if( (sockfd = socket(AF_INET, SOCK_DGRAM, 0)) <0){
        perror("socket creation failed");
        exit(EXIT_FAILURE);
    }

    memset(&servaddr, 0, sizeof(servaddr));

    // Filling server information
    servaddr.sin_family = AF_INET;
    servaddr.sin_port = htons(PORT);
    servaddr.sin_addr.s_addr = INADDR_ANY;

    ssize_t bytes_received,bytes_sent;
    bytes_sent = sendto(sockfd, hello, strlen(hello), 0,(struct sockaddr*)&servaddr,sizeof(servaddr));
    if(bytes_sent < 0){
        perror("send to failed");
        close(sockfd);
        exit(EXIT_FAILURE);
    } 

    printf("Sent to server : %s\n",hello);

    socklen_t len = sizeof(servaddr);
    bytes_received = recvfrom(sockfd, buffer, BUFFER_SIZE-1, 0, (struct sockaddr*)&servaddr, &len);
    if(bytes_received < 0){
        perror("recvfrom failed");
        close(sockfd);
        exit(EXIT_FAILURE);
    }
    buffer[bytes_received] = '\0';
    printf("Receive from server: %s\n",buffer);
   
    close(sockfd);

    return 0;
}