/*Common setup code(Common to Both mode ET and LT)*/ 
#include <string.h>
#include<stdio.h>
#include<fcntl.h>
#include<sys/socket.h>
#include<unistd.h>
#define MAX_EVENTS 10
#define READ_BUF_SIZE 1024

// Helper function to set non-blocking mode
int set_nonblocking(int fd){
  int flags = fcntl(fd,F_GETFL,0);
  if(flags == -1) return -1;
  return fcntl(fd, F_SETFL, flags | O_NONBLOCK);
}

// Create and Configure listening socket
int create_listening_socket(int port){
  int listen_fd = socket(AF_INET,SOCK_STREAM,0);
  if (listen_fd == -1) return -1;


  // Set SO_REUSEADDR to avoid "address already in use" erros
  int reuse = 1;
  setsockopt(listen_fd, SOL_SOCKET, SO_REUSEADDR, &reuse , sizeof(reuse));

  struct sockaddr_in addr;
  memset(&addr, 0, sizeof(addr));
  // TODO:sett  addr and bind fd.
}
