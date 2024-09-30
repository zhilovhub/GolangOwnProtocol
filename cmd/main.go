package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Server
	go func() {
		listener, err := net.Listen("tcp", ":8080")
		if err != nil {
			fmt.Printf("Error while listening: %v", err)
			return
		}
		defer listener.Close()

		for {
			conn, err := listener.Accept()
			fmt.Print(1)
			if err != nil {
				fmt.Printf("Error while accepting connection: %v", err)
				continue
			}
			handleConnection(conn)
		}
	}()

	// Client
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Error connectiong: %v", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("Привте"))
	conn.Write([]byte("Как там дела?"))
	time.Sleep(time.Second * 10)
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println("Got:", string(buffer))
}
