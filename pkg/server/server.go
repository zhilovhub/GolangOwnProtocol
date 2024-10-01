package server

import (
	"bufio"
	"fmt"
	"net"
)

func Server() {
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
			fmt.Printf("Error while accepting connection: %v\n", err)
			continue
		}
		go func(conn net.Conn) {
			reader := bufio.NewReader(conn)
			for {
				bytes, _, err := reader.ReadLine()
				if err != nil {
					fmt.Printf("Error while reading bytes from connection: %v\n", err)
					return
				}
				fmt.Printf("Got new message with %d bytes: %s\n", len(bytes), string(bytes))
			}
		}(conn)
	}
}
