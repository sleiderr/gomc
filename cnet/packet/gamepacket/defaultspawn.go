package gamepacket

import (
	"bytes"
	"encoding/binary"

	"github.com/sleiderr/gomc/ctypes"
	"github.com/sleiderr/gomc/game"
)

type DefaultSpawnPosition struct {
	Location game.EntityLocation
	Angle    float32
}

func (p *DefaultSpawnPosition) Raw() []byte {
	rawData := new(bytes.Buffer)

	ctypes.WritePosition(rawData, p.Location)
	binary.Write(rawData, binary.BigEndian, p.Angle)

	return rawData.Bytes()
}

func (p *DefaultSpawnPosition) FillFromRaw(raw []byte) error {
	return nil
}
