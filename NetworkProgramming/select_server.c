#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#define PORT 8080
#define BUFFER_SIZE 1024
#define MAX_SIZE 1024
int main(){
    int sockfd;
    char buffer[BUFFER_SIZE];
    struct sockaddr_in serveaddr;

    if((sockfd = socket(AF_INET, SOCK_STREAM, 0))<0){
        perror("socket creation failed");
        exit(EXIT_FAILURE);
    }

    memset(&serveaddr, 0, sizeof(serveaddr));

    serveaddr.sin_family = AF_INET;
    serveaddr.sin_addr.s_addr = INADDR_ANY;
    serveaddr.sin_port = htons(PORT);

    if(bind(sockfd, (struct sockaddr*)&serveaddr, sizeof(serveaddr))<0){
        perror("Bind failed");
        exit(EXIT_FAILURE);
    }

    if(listen(sockfd, 5)<0){
        perror("listen");
        exit(EXIT_FAILURE);
    }

    





    return 0;
}