package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func main() {
	// Server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error while listening: %v", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error while accepting connection: %v", err)
			continue
		}
		go func(conn net.Conn) {
			var buffer bytes.Buffer
			_, err := io.Copy(&buffer, conn)
			if err != nil {
				fmt.Printf("Error while reading bytes from connection: %v", err)
				return
			}
			fmt.Printf("Got new message: %s\n", buffer.String())
		}(conn)
	}
}
