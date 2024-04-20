package gamepacket

import (
	"bytes"
	"encoding/binary"

	"github.com/sleiderr/gomc/ctypes"
	"github.com/sleiderr/gomc/game"
	"github.com/sleiderr/gomc/utils"
)

type PlayLoginPacket struct {
	EntityID            int32
	IsHardcore          bool
	DimensionCount      int32
	DimensionNames      []string
	MaxPlayers          int32
	ViewDistance        int32
	SimulationDistance  int32
	ReducedDebug        bool
	EnableRespawnScreen bool
	DoLimitedCrafting   bool
	DimensionType       string
	DimensionName       string
	HashedSeed          uint64
	GameMode            byte
	PreviousGameMode    byte
	IsDebug             bool
	IsFlat              bool
	HasDeathLocation    bool
	DeathDimensionName  string
	DeathLocation       game.EntityPosition
	PortalCooldown      int32
}

func (p *PlayLoginPacket) Raw() []byte {
	rawData := new(bytes.Buffer)
	binary.Write(rawData, binary.BigEndian, p.EntityID)
	rawData.WriteByte(utils.EncodeBool(p.IsHardcore))
	ctypes.AsVarInt(p.DimensionCount).WriteVarInt(rawData)
	for _, dim := range p.DimensionNames {
		ctypes.WriteString(rawData, dim)
	}
	ctypes.AsVarInt(p.MaxPlayers).WriteVarInt(rawData)
	ctypes.AsVarInt(p.ViewDistance).WriteVarInt(rawData)
	ctypes.AsVarInt(p.SimulationDistance).WriteVarInt(rawData)
	rawData.WriteByte(utils.EncodeBool(p.ReducedDebug))
	rawData.WriteByte(utils.EncodeBool(p.EnableRespawnScreen))
	rawData.WriteByte(utils.EncodeBool(p.DoLimitedCrafting))
	ctypes.WriteString(rawData, p.DimensionType)
	ctypes.WriteString(rawData, p.DimensionName)
	binary.Write(rawData, binary.BigEndian, p.HashedSeed)
	rawData.WriteByte(p.GameMode)
	rawData.WriteByte(p.PreviousGameMode)
	rawData.WriteByte(utils.EncodeBool(p.IsDebug))
	rawData.WriteByte(utils.EncodeBool(p.IsFlat))
	rawData.WriteByte(utils.EncodeBool(p.HasDeathLocation))
	if p.HasDeathLocation {
		ctypes.WriteString(rawData, p.DeathDimensionName)
	}
	ctypes.AsVarInt(p.PortalCooldown).WriteVarInt(rawData)

	return rawData.Bytes()
}

func (p *PlayLoginPacket) FillFromRaw(raw []byte) error {
	return nil
}
