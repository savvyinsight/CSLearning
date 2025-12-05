#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/types.h>

#define PORT 8080
#define BUFFER_SIZE 1024

int main(){
    int sockfd;
    struct sockaddr_in servaddr,cliaddr;
    char buffer[BUFFER_SIZE];

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
    

    return 0;
}