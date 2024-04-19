package packet

import (
	"github.com/sleiderr/gomc/ctypes"
)

type CraftPacket struct {
	packetId CraftPacketType
	payload  CraftPacketPayload
}

type CraftPacketPayload interface {
	Raw() []byte
	FillFromRaw([]byte) error
}

func NewCraftPacket(packetId CraftPacketType, payload CraftPacketPayload) *CraftPacket {
	return &CraftPacket{
		packetId,
		payload,
	}
}

func ParseRaw(raw *RawCraftPacket, state byte) (*CraftPacket, error) {
	var payload CraftPacketPayload

	packetType := CraftPacketType(NewPacketType(state, uint32(raw.packetId.Value())))
	payload, err := MakePacket(packetType)

	if err != nil {
		return nil, err
	}

	err = payload.FillFromRaw(raw.Data())

	if err != nil {
		return nil, err
	}

	return &CraftPacket{
		packetId: packetType,
		payload:  payload,
	}, nil
}

func (p *CraftPacket) AsRaw() *RawCraftPacket {
	return &RawCraftPacket{
		packetId: ctypes.AsVarInt(int32(p.packetId.Id)),
		data:     p.payload.Raw(),
	}
}

func (p *CraftPacket) Id() CraftPacketType {
	return p.packetId
}

func (p *CraftPacket) Payload() CraftPacketPayload {
	return p.payload
}
