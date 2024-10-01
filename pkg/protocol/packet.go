package protocol

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
