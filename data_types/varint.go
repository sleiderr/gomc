package data_types

import (
	"bufio"
	"errors"
	"io"
)

type VarInt struct {
	val int32
}

func ReadVarInt(read io.Reader) (VarInt, error) {
	buff := bufio.NewReader(read)
	var val int32
	pos := 0

	for {
		currByte, err := buff.ReadByte()

		if err != nil {
			return VarInt{}, err
		}

		val |= (int32(currByte) & 0x7F) << pos

		if (currByte & 0x80) == 0 {
			break
		}

		pos += 7

		if pos >= 32 {
			return VarInt{}, errors.New("Invalid VarInt (too long)")
		}
	}

	return VarInt{val}, nil
}

func (varint *VarInt) WriteVarInt(w io.Writer) {
	bufWr := bufio.NewWriter(w)
	val := varint.val
	for {
		if (val & 0x80) == 0 {
			bufWr.WriteByte(byte(val))
			return
		}

		bufWr.WriteByte((byte(val) & 0x7F) | 0x80)

		val = int32(uint32(val) >> 7)
	}
}
