package packet

import (
	"fmt"
	"reflect"

	"github.com/sleiderr/gomc/cnet/packet/config"
	"github.com/sleiderr/gomc/cnet/packet/gamepacket"
	"github.com/sleiderr/gomc/cnet/packet/login"
	"github.com/sleiderr/gomc/cnet/packet/slp"
	"github.com/sleiderr/gomc/utils"
)

const (
	HandshakePacketId    = 0
	LoginStartPacketId   = 0
	ClientConfigPacketId = 0
	PingPacketId         = 1
	ConfigPluginMessage  = 1
	LoginSuccessPacketId = 2
	FinishConfigPacketId = 2
	LoginAckPacketId     = 3
	PlayLoginPacketId    = 0x29
	PlaySyncPosition     = 0x3E
	PlaySetSpawnPos      = 0x54
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
	registerPacketType(NewPacketType(2, 0), (*login.LoginStart)(nil))
	registerPacketType(NewPacketType(2, 3), (*login.LoginAck)(nil))
	registerPacketType(NewPacketType(3, 0), (*config.ClientInformation)(nil))
	registerPacketType(NewPacketType(3, ConfigPluginMessage), (*config.PluginMessage)(nil))
	registerPacketType(NewPacketType(3, FinishConfigPacketId), (*utils.FieldlessPacket)(nil))
	registerPacketType(NewPacketType(4, PlayLoginPacketId), (*gamepacket.PlayLoginPacket)(nil))
}

func MakePacket(t CraftPacketType) (CraftPacketPayload, error) {
	fmt.Println(t)
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
