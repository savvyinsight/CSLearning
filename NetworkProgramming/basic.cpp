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

S:socket, bind, listen, accept, send/recv, close
C:socket                connect,send/recv, close               
You can know all api details by checking.

---------------UDP Newtork Programming API------------------
socket, bind, sendto, recvfrom, close

---------------convert IPv4 and IPv6 addresses--------------
inet_pton , inet_ntop :  convert IPv4 and IPv6 addresses from text to binary form
inet_addr :              converts the Internet host address from the IPv4 numbers-and-dots notation into binary form (in network  byte
order)

---------------I/O Multiplexing--------------
I/O Multiplexing allows a single process or thread to monitor multiple input/output (I/O) 
streams — such as sockets, files, or pipes —
simultaneously, without blocking on any one of them. Instead of using one thread per I/O source,
multiplexing enables efficient handling of numerous I/O operations through a single control 
loop, typically via system calls like select(), poll(), or epoll() in Unix-like operating systems.

1.select():
    FD_ZERO(fd_set *set)     // Clear all bits in set
    FD_SET(int fd, fd_set *set)   // Add fd to set
    FD_CLR(int fd, fd_set *set)   // Remove fd from set
    FD_ISSET(int fd, fd_set *set) // Test if fd is in set
2.poll()
3.epoll():epoll_create, epoll_create1
    epoll_ctl
    epoll_wait, epoll_pwait

-----------------------Blocking and Non-Blocking---------------
Blocking : blocking sockets is default mode, when call 'recv(),send(),accept()',the function blocks
(pause execution) the thread or progress until:
For 'recv()': Data is available from network.
For 'send()': Data is abailable from kernel's send buffer.
For 'accept()': A new client connection is pending.

Non-Blocking: immediately return. these 'recv(),send(),accept()'will be set to 'EAGAIN' or 'EWOULDBLOCK'
key advandtage : let programmes perform other tasks while waiting for i/o,enample handiling of 
multiple connections with i/o multiplexing.

How to do?
To set a socket to Non-Blocking mode , we use the 'fcntl()'(file control) system call.
see details with man fcntl 

===>Converting a TCP Socket to non-block:
1.create tcp socket : int sockfd = socket(AF_INET,SOCK_STREAM,0)
2.get current socket flags : int flags = fcntl(sockfd,F_GETFL)
3.add non-block to flag: flags |= O_NONBLOCK
4.apply the new flag: // Apply new flags
if (fcntl(sockfd, F_SETFL, flags) < 0) {
    perror("fcntl F_SETFL failed");
    ...
}

*/

