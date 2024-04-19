package packet

import (
	"bytes"
	"io"

	"github.com/sleiderr/gomc/ctypes"
)

type RawCraftPacket struct {
	packetId ctypes.VarInt
	data     []byte
}

func ReadCraftPacket(r io.Reader) (*RawCraftPacket, error) {
	varLength, err := ctypes.ReadVarInt(r)
	packetId, err := ctypes.ReadVarInt(r)

	if err != nil {
		return nil, err
	}

	packetLength := int(varLength.Value()) - packetId.Length()
	packetData := make([]byte, packetLength)

	bytesRead, err := io.ReadFull(r, packetData)

	if err != nil {
		return nil, err
	}

	if bytesRead < packetLength {
		return nil, err
	}

	return &RawCraftPacket{
		packetId: packetId,
		data:     packetData,
	}, nil
}

func (rP *RawCraftPacket) Bytes() []byte {
	buf := new(bytes.Buffer)
	rP.packetId.WriteVarInt(buf)

	rawBytes := new(bytes.Buffer)
	ctypes.AsVarInt(int32(buf.Len()) + int32(len(rP.data))).WriteVarInt(rawBytes)
	rP.packetId.WriteVarInt(rawBytes)
	rawBytes.Write(rP.data)

	return rawBytes.Bytes()
}

func (pack *RawCraftPacket) Data() []byte {
	return pack.data
}
