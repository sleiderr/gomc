package gamepacket

import (
	"bytes"
	"encoding/binary"

	"github.com/sleiderr/gomc/ctypes"
	"github.com/sleiderr/gomc/game"
)

type TeleportFlags byte

type SynchronizePositionPacket struct {
	Position   game.EntityPosition
	Flags      TeleportFlags
	TeleportID int32
}

func (p *SynchronizePositionPacket) Raw() []byte {
	rawData := new(bytes.Buffer)

	binary.Write(rawData, binary.BigEndian, p.Position.X)
	binary.Write(rawData, binary.BigEndian, p.Position.Y)
	binary.Write(rawData, binary.BigEndian, p.Position.Z)
	binary.Write(rawData, binary.BigEndian, p.Position.Yaw)
	binary.Write(rawData, binary.BigEndian, p.Position.Pitch)
	rawData.WriteByte(byte(p.Flags))
	ctypes.AsVarInt(p.TeleportID).WriteVarInt(rawData)

	return rawData.Bytes()
}

func (p *SynchronizePositionPacket) FillFromRaw(raw []byte) error {
	return nil
}
