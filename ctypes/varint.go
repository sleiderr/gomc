package ctypes

import (
	"bufio"
	"errors"
	"io"
)

type VarInt struct {
	val    int32
	length int
}

func ReadVarInt(r io.Reader) (VarInt, error) {
	var val int32
	pos, length := 0, 0
	var currByte [1]byte

	for {
		_, err := r.Read(currByte[:])
		length += 1

		if err != nil {
			return VarInt{}, err
		}

		val |= (int32(currByte[0]) & 0x7F) << pos

		if (currByte[0] & 0x80) == 0 {
			break
		}

		pos += 7

		if pos >= 32 {
			return VarInt{}, errors.New("Invalid VarInt (too long)")
		}
	}

	return VarInt{val, length}, nil
}

func AsVarInt(val int32) VarInt {
	return VarInt{val, 4}
}

func (varint VarInt) Value() int32 {
	return varint.val
}

func (varint VarInt) Length() int {
	return varint.length
}

func (varint VarInt) WriteVarInt(w io.Writer) {
	bufWr := bufio.NewWriter(w)
	val := varint.val
	for {
		if (val & 0x80) == 0 {
			bufWr.WriteByte(byte(val))
			bufWr.Flush()
			return
		}

		bufWr.WriteByte((byte(val) & 0x7F) | 0x80)

		val = int32(uint32(val) >> 7)
	}
}
