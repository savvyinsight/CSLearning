#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <unistd.h>
#include <arpa/inet.h>

#define PORT 8080
#define BUFFER_SIZE 1024

int main(){
    int sockfd;
    struct sockaddr_in servaddr,cliaddr;
    char buffer[BUFFER_SIZE];
    const char* hello = "Hello From Server";

    sockfd = socket(AF_INET, SOCK_DGRAM, 0);
    if(sockfd < -1){
        perror("socket creation failed");
        exit(EXIT_FAILURE);
    }

    memset(&servaddr, 0, sizeof(servaddr));

    servaddr.sin_family = AF_INET;
    servaddr.sin_addr.s_addr = INADDR_ANY;
    servaddr.sin_port = htons(PORT);

    if(bind(sockfd,(const struct sockaddr*)&servaddr,sizeof(servaddr))<0){
        perror("bind failed.");
        exit(EXIT_FAILURE);
    }

    printf("UDP listening...\n");

    ssize_t bytes_received,bytes_sent;
    socklen_t addr_len = sizeof(cliaddr);
    while (1) {
        bytes_received = recvfrom(sockfd, buffer, BUFFER_SIZE-1, 0, (struct sockaddr*)&cliaddr,&addr_len);
        if(bytes_received < 0){
            perror("receive failed.");
            continue;
        }

        buffer[bytes_received] = '\0';//Null-termination
        char *ip_str = inet_ntoa(cliaddr.sin_addr);
        int port = ntohs(cliaddr.sin_port);
        printf("Client ip:%s  port:%d, message:%s\n",ip_str,port,buffer);

        bytes_sent = sendto(sockfd, hello, strlen(hello), 0, (struct sockaddr*)&cliaddr, addr_len);
        if(bytes_sent < 0){
            perror("sent  failed.");
            continue;
        }
        printf("hello message sent.\n");

    }
    close(sockfd);

    return 0;
}