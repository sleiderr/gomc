package packet

import (
	"fmt"

	"github.com/sleiderr/gomc/cnet/packet/slp"
	"github.com/sleiderr/gomc/ctypes"
)

type CraftPacketType int32

const (
	Handshake CraftPacketType = 0
	Ping      CraftPacketType = 1
)

type CraftPacket struct {
	packetId CraftPacketType
	payload  CraftPacketPayload
}

type CraftPacketPayload interface {
	Raw() []byte
}

func NewCraftPacket(packetId CraftPacketType, payload CraftPacketPayload) *CraftPacket {
	return &CraftPacket{
		packetId,
		payload,
	}
}

func ParseRaw(raw *RawCraftPacket) (*CraftPacket, error) {
	var payload CraftPacketPayload
	var err error
	switch packetType := CraftPacketType(raw.packetId.Value()); packetType {
	case Handshake:
		p := &HandshakePacket{}
		err = p.FillFromRaw(raw.Data())
		payload = p
	case Ping:
		p := &slp.PingPacket{}
		err = p.FillFromRaw(raw.Data())
		payload = p
	default:
		return nil, fmt.Errorf("Unexpected packet")
	}

	if err != nil {
		return nil, err
	}

	return &CraftPacket{
		packetId: CraftPacketType(raw.packetId.Value()),
		payload:  payload,
	}, nil
}

func (p *CraftPacket) AsRaw() *RawCraftPacket {
	return &RawCraftPacket{
		packetId: ctypes.AsVarInt(int32(p.packetId)),
		data:     p.payload.Raw(),
	}
}

func (p *CraftPacket) Id() CraftPacketType {
	return p.packetId
}

func (p *CraftPacket) Payload() CraftPacketPayload {
	return p.payload
}
