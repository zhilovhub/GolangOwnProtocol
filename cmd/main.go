package main

import (
	"fmt"
	"version1/pkg/protocol"
)

func main() {
	packet := protocol.CreateIPacket(2, 3)
	fmt.Println(packet.Fields)
}
