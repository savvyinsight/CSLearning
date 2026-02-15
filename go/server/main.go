package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		n, err := conn.Read(buf)
		if err != nil {
			return
		}

		fmt.Println(string(buf[:n]))
	}
}

func main() {
	fmt.Println("Listen....")

	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("listen err")
		return
	}

	for {
		conn, err1 := listen.Accept()
		if err1 != nil {
			fmt.Println("connect faild")
			return
		} else {
			fmt.Printf("connected:%v, client addr:%d\n", conn, conn.RemoteAddr().String())
		}

		go process(conn)

	}
}
