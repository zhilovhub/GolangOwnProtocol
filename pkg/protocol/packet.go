package protocol

import (
	"bytes"
	"slices"
)

type IPacket struct {
	PacketType    byte
	PacketSubType byte
	Fields        []IPacketField
}

type IPacketField struct {
	FieldId       byte
	FieldSize     byte
	FieldContents []byte
}

func CreateIPacket(packetType, packetSubType byte) *IPacket {
	return &IPacket{
		PacketType:    packetType,
		PacketSubType: packetSubType,
	}
}

func (p *IPacket) ToPacket() ([]byte, error) {
	buffer := new(bytes.Buffer)
	_, err := buffer.Write([]byte{0xDD, 0xEF, 0xDD, p.PacketType, p.PacketSubType})
	if err != nil {
		return nil, err
	}

	slices.SortFunc(p.Fields, func(a, b IPacketField) int {
		if a.FieldId < b.FieldId {
			return -1
		}
		if a.FieldId == b.FieldId {
			return 0
		}
		return 1
	})

	for _, field := range p.Fields {
		buffer.Write([]byte{field.FieldId, field.FieldSize})
		buffer.Write(field.FieldContents)
	}

	buffer.Write([]byte{0x00, 0xFF})
	return buffer.Bytes(), nil
}
