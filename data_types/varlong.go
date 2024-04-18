package data_types

import (
	"bufio"
	"errors"
	"io"
)

type VarLong struct {
	val int64
}

func ReadVarLong(r io.Reader) (VarLong, error) {
	buf := bufio.NewReader(r)
	var val int64
	pos := 0

	for {
		currByte, err := buf.ReadByte()

		if err != nil {
			return VarLong{}, nil
		}

		val |= (int64(currByte) & 0x7F) << pos

		if (currByte & 0x80) == 0 {
			break
		}

		pos += 7

		if pos >= 32 {
			return VarLong{}, errors.New("Unexpected byte read (VarLong)")
		}
	}

	return VarLong{val}, nil
}

func (varl *VarLong) WriteVarLong(w io.Writer) {
	val := varl.val
	buf := bufio.NewWriter(w)

	for {
		if (val & 0x80) == 0 {
			buf.WriteByte(byte(val & 0xFF))
			return
		}

		buf.WriteByte((byte(val&0xFF) & 0x7F) | 0x80)

		val = int64(uint64(val) >> 7)
	}
}
