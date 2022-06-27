package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "localhost:2000")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err)
		os.Exit(1)
	}
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			return
		}
		fmt.Print("Server: ", message)
		conn.Write([]byte("msg from socket " + message + "\n"))
	}
}
