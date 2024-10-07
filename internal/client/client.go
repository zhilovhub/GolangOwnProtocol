package client

import (
	"fmt"
	"net"
	"time"
)

func Client() {
	// Client
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Error connectiong: %v", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("Привте"))
	time.Sleep(time.Second * 2)
}