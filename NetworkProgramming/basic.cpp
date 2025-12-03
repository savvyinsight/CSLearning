/*
--------------About Socket ------------

man soket  ->to see details
1.Socket
    - create an endpoint for communication
    type:
    - SOCK_STREAM ->TCP based
    - SOCK_DGRAM ->UDP based

2.S/C model
    - Server and Client

3.socket API.

--------------About TCP S/C Communication Process ------------
1.Server Side
    1.Create Socket : int socket(int domain, int type, int protocol);
    2.Bind : int bind(int sockfd, const struct sockaddr *addr,
                socklen_t addrlen);

    ''' Note: Details look at by man <..>
    3.Listen
    4.Accept
    5.Send and Receive
    6.Close 


2.Client Side
    1.Create Socket
    2.Connect
    3.Send and Receive
    4.Close

--------------TCP Connection Establishment Process :3 way HandShake-------------------------------
Refer : https://ipwithease.com/understanding-tcp-3-way-handshake-process/
Refer: https://www.geeksforgeeks.org/computer-networks/tcp-3-way-handshake-process/

==========Send and Receive data ===================

-------------- TCP Connection Termination : 4 way four-step handshake-------------------------
Refer : https://www.geeksforgeeks.org/computer-networks/why-tcp-connect-termination-need-4-way-handshake/




---------------TCP Newtork Programming API------------------
You Can check a specific api by : man 7 socket   or   man socket,  
if you want to know in which chapter : man -k socket

You can know all api details by checking.
*/


#include <cstddef>
#include <cstring>
#include <iostream>
#include <ostream>
#include <string>
#include <sys/types.h>
int main(){
    std::cout<<sizeof(int)<<std::endl;
    std::cout<<sizeof(size_t)<<std::endl;
    std::cout<<sizeof(ssize_t)<<std::endl;
    std::cout<<sizeof(long)<<std::endl;


    

}