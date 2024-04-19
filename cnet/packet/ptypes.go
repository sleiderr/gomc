package packet

import (
	"fmt"
	"reflect"

	"github.com/sleiderr/gomc/cnet/packet/slp"
)

const (
	HandshakePacketId = 0
	PingPacketId      = 1
)

type CraftPacketType struct {
	State byte
	Id    uint32
}

func NewPacketType(state byte, id uint32) CraftPacketType {
	return CraftPacketType{state, id}
}

var packetTypeRegistry = make(map[CraftPacketType]reflect.Type)

func InitPacketTypes() {
	registerPacketType(NewPacketType(0, 0), (*HandshakePacket)(nil))
	registerPacketType(NewPacketType(1, 0), (*HandshakePacket)(nil))
	registerPacketType(NewPacketType(1, 1), (*slp.PingPacket)(nil))
}

func MakePacket(t CraftPacketType) (CraftPacketPayload, error) {
	if _, ok := packetTypeRegistry[t]; !ok {
		return nil, fmt.Errorf("Unexpected packet (invalid packetId and state combination)")
	}

	p, ok := reflect.New(packetTypeRegistry[t].Elem()).Interface().(CraftPacketPayload)

	if !ok {
		return nil, fmt.Errorf("Unexpected packet (invalid packetId and state combination)")
	}

	return p, nil
}

func registerPacketType(id CraftPacketType, pTypeNil CraftPacketPayload) {
	packetTypeRegistry[id] = reflect.TypeOf(pTypeNil)
}
