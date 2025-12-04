# issues
## Writing basic tcp server/client
### 1.Buffer Handling Issues
Always writing/reading the full BUFFER_SIZE (1024 bytes), even if the actual message is shorter. This wastes bandwidth and can cause problems.

**Better approach:** Track the actual message length and use that:

```C
// In client.c chat() function:
size_t len = strlen(buffer);
write(sockfd, buffer, len);  // Send only actual data
read(sockfd, buffer, BUFFER_SIZE - 1);  // Leave room for null terminator
```

### 2.Unchecked Return Values
Check the return values of read() and write(). These calls can fail or read partial data.
```C
ssize_t bytes_read = read(sockfd, buffer, BUFFER_SIZE - 1);
if (bytes_read <= 0) {
    // Handle error or connection closed
    break;
}
buffer[bytes_read] = '\0';  // Null-terminate
```
### 3.Exit Condition Logic
Exit condition checks strncmp(buffer, "exit", 4), but this could match messages starting with "exit" anywhere in the buffer.
**Better:**
```C
// Check if message is exactly "exit\n"
if (strcmp(buffer, "exit\n") == 0) {
    printf("Exiting...\n");
    break;
}
```

### 4.Minor Issues
- #include <strings.h> for bzero() - prefer memset() for portability
- Server doesn't close connfd after chat ends

### Some Questions
#### bzero and memset, which is better and why?
memset() is better and should be preferred.

**Why:**

**Portability**: memset() is standard C/C++. bzero() is **deprecated** and removed from POSIX standards.

**Availability**: memset() works everywhere. bzero() is mostly Linux/UNIX specific.

**Clarity**: memset(buffer, 0, size) clearly shows you're setting to zero vs bzero(buffer, size).

Equivalent usage:
```C
bzero(buffer, sizeof(buffer));          // Deprecated
memset(buffer, 0, sizeof(buffer));      // Standard, portable
```

#### send() and recv() instead of read()/write() for sockets,why?
send()/recv() are better for sockets.

**Why**:

1.**Socket-specific flags:** send()/recv() have extra flags like MSG_NOSIGNAL, MSG_DONTWAIT

2.**Clear intent:** Shows you're specifically working with sockets, not generic file I/O

3.**Same functionality:** They work identically to read()/write() for basic usage
```C
// With flags (can't do this with write/read)
send(sockfd, buffer, len, MSG_NOSIGNAL);  // No SIGPIPE signal
recv(sockfd, buffer, size, MSG_DONTWAIT); // Non-blocking
```
