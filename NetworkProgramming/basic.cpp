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
    
*/ 