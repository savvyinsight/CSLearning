#include <fcntl.h>
#include <netinet/in.h>
#include <stdio.h>
#include <sys/socket.h>
int set_nonblocking(int sockfd){
    int flags = fcntl(sockfd,F_GETFL,0);
    if(flags<0){
        perror("fnctl F_GETFL");
        return -1;
    }

    flags |= O_NONBLOCK;

    if(fcntl(sockfd, F_SETFL,flags)<0){
        perror("fcntl F_SETFL");
        return -1;
    }
    return 0;
}

int main(){
    int server_fd,client_fd;
    struct sockaddr_in serveraddr,client_addr;


    return 0;
}

