package main

import (
	"fmt"
	"log"
	"version1/pkg/protocol"
)

func main() {
	packet := protocol.CreateIPacket(2, 3)
	packetBytes, err := packet.ToPacket()
	if err != nil {
		log.Fatalf("Error while converting packet to bytes: %v\n", err)
	}

	fmt.Println(packetBytes)
}
