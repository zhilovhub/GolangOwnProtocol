package protocol

import (
	"bytes"
	"fmt"
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

// Creates IPacket with parameters
func CreateIPacket(packetType, packetSubType byte) *IPacket {
	return &IPacket{
		PacketType:    packetType,
		PacketSubType: packetSubType,
	}
}

// Parses byte array to IPacket structure
func ParsePacket(b []byte) (*IPacket, error) {
	bytesLength := len(b)
	if bytesLength < 7 {
		return nil, fmt.Errorf("bytes length is less than 7 bytes: %d", bytesLength)
	}

	if b[0] != 0xDD || b[1] != 0xEF || b[2] != 0xDD {
		return nil, fmt.Errorf("bytes has wrong headers")
	}

	var endIndex = bytesLength - 1
	if b[endIndex-1] != 0x00 || b[endIndex] != 0xFF {
		return nil, fmt.Errorf("bytes has wrong end bytes")
	}

	packetType := b[3]
	packetSubType := b[4]

	iPacket := CreateIPacket(packetType, packetSubType)

	fieldsBytes := b[5:]
	for {
		if len(fieldsBytes) == 2 {
			break
		}
		fieldId := fieldsBytes[0]
		fieldSize := fieldsBytes[1]

		var fieldContents []byte
		if fieldSize != 0 {
			fieldContents = fieldsBytes[:fieldSize]
		}

		iPacket.Fields = append(iPacket.Fields, IPacketField{
			FieldId:       fieldId,
			FieldSize:     fieldSize,
			FieldContents: fieldContents,
		})
		fieldsBytes = fieldContents[2+fieldSize:]
	}
	return iPacket, nil
}

// Convers IPacket to Packet in byte slice presentation
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
