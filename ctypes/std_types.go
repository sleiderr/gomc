package ctypes

import (
	"fmt"
	"io"
	"unicode/utf8"
)

func WriteString(w io.Writer, data string) {
	strLength := 0

	for _, ch := range data {
		if utf8.RuneLen(ch) > 3 {
			strLength += 1
		}
		strLength += 1
	}

	AsVarInt(int32(strLength)).WriteVarInt(w)
	w.Write([]byte(data))
}

func ReadString(r io.Reader) (string, error) {
	varlength, err := ReadVarInt(r)
	val := ""

	if err != nil {
		return "", err
	}

	length := varlength.val

	for length > 0 {
		buf := make([]byte, 1)
		_, err := r.Read(buf)
		if err != nil {
			return "", err
		}

		firstByte := buf[0]
		runeLength := utf8.RuneLen(rune(firstByte))
		if runeLength <= 0 {
			return "", fmt.Errorf("invalid UTF-8 encoding")
		}

		buf = make([]byte, runeLength-1)
		_, err = r.Read(buf)
		if err != nil {
			return "", err
		}

		combinedBytes := append([]byte{firstByte}, buf...)
		currCodePoint, _ := utf8.DecodeRune(combinedBytes)

		if runeLength > 3 {
			length -= 1
		}

		length -= 1
		val += string(currCodePoint)
	}

	return val, nil
}
