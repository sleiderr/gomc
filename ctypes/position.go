package ctypes

import (
	"encoding/binary"
	"io"

	"github.com/sleiderr/gomc/game"
)

func ReadPosition(r io.Reader) (game.EntityLocation, error) {
	var loc int64
	err := binary.Read(r, binary.BigEndian, &loc)

	if err != nil {
		return game.EntityLocation{}, err
	}

	x := loc >> 38
	y := loc << 52 >> 52
	z := loc << 26 >> 38

	return game.EntityLocation{
		X: int32(x),
		Y: int16(y),
		Z: int32(z),
	}, nil
}

func WritePosition(w io.Writer, pos game.EntityLocation) error {
	encodedPos := ((int64(pos.X) & 0x3FFFFFF) << 38) | ((int64(pos.Z) & 0x3FFFFFF) << 12) | (int64(pos.Y) & 0xFFF)

	return binary.Write(w, binary.BigEndian, encodedPos)
}
