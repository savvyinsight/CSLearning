package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Client...")

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("failed to connnect")
		return
	}

	fmt.Println("connect success", conn)

	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("failed...")
		return
	}

	n, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println("conn fail")
	}

	fmt.Println("fa", n)
	conn.Close()
}
