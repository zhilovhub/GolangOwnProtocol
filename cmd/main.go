package main

import (
	"fmt"
	"version1/pkg/protocol"
)

func main() {
	iPacket := protocol.CreateIPacket(1, 1)
	protocol.SetValue(iPacket, 1, int32(2))

	byteArray, err := iPacket.ToPacket()
	if err != nil {
		panic(err)
	}

	newIPacket, err := protocol.ParsePacket(byteArray)
	if err != nil {
		panic(err)
	}

	var value interface{}
	value, err = protocol.GetValue[int32](newIPacket, 1)
	if err != nil {
		panic(err)
	}

	if v, ok := value.(int32); ok {
		fmt.Println(v)
	} else {
		panic(fmt.Errorf("value has another type: %T", v))
	}
}
