package main

import (
	"fmt"
	"version1/internal/protocol"
)

func main() {
	iPacket := protocol.CreateIPacket(1, 1)
	protocol.SetValue(iPacket, 1, int64(34573457))
	protocol.SetValue(iPacket, 2, float32(23.43))
	protocol.SetValue(iPacket, 3, float64(34573457.12))
	protocol.SetValue(iPacket, 4, true)

	byteArray, err := iPacket.ToPacket()
	if err != nil {
		panic(err)
	}

	newIPacket, err := protocol.ParsePacket(byteArray)
	if err != nil {
		panic(err)
	}

	value, err := protocol.GetValue[int64](newIPacket, 1)
	if err != nil {
		panic(err)
	}
	value1, err := protocol.GetValue[float32](newIPacket, 2)
	if err != nil {
		panic(err)
	}
	value2, err := protocol.GetValue[float64](newIPacket, 3)
	if err != nil {
		panic(err)
	}
	value3, err := protocol.GetValue[bool](newIPacket, 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(value, value1, value2, value3)

}
