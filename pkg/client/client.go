package main

import (
	"fmt"
	"net"
)

func main() {
	// Client
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Error connectiong: %v", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("Привте"))
}
