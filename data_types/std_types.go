package data_types

import (
	"bufio"
	"io"
)

func ReadString(r io.Reader) (string, error) {
	varlength, err := ReadVarInt(r)
	buf := bufio.NewReader(r)
	val := ""

	if err != nil {
		return "", err
	}

	length := varlength.val

	for length > 0 {
		currCodePoint, size, err := buf.ReadRune()

		if err != nil {
			return "", err
		}

		if size > 3 {
			length -= 1
		}

		length -= 1
		val += string(currCodePoint)
	}

	return val, nil
}
