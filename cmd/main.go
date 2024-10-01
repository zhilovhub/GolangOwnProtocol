package main

import (
	"fmt"
	"log"
	"version1/pkg/protocol"
)

func main() {
	iPacket := protocol.CreateIPacket(2, 3)
	packetBytes, err := iPacket.ToPacket()
	if err != nil {
		log.Fatalf("Error while converting IPacket to bytes: %v\n", err)
	}

	iPacket, err = protocol.ParsePacket(packetBytes)
	if err != nil {
		log.Fatalf("Error while converting bytes to IPacket: %v\n", err)
	}

	fmt.Printf("Packet: %v\n", packetBytes)
	fmt.Printf("IPacket: %v\n", iPacket)

}
