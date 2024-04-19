package packet

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/sleiderr/gomc/ctypes"
)

type HandshakePacket struct {
	StatusReq    bool
	ProtoVersion int32
	ServAddr     string
	ServPort     uint16
	NextState    int32
}

type StatusResponsePacket struct {
	Status string
}

func (p *StatusResponsePacket) FillFromRaw(raw []byte) error {
	return nil
}

func (p *StatusResponsePacket) Raw() []byte {
	rawData := new(bytes.Buffer)
	ctypes.WriteString(rawData, p.Status)

	return rawData.Bytes()
}

func (p *HandshakePacket) Raw() []byte {
	if p.StatusReq {
		return make([]byte, 0)
	}

	rawData := new(bytes.Buffer)

	ctypes.AsVarInt(p.ProtoVersion).WriteVarInt(rawData)
	ctypes.WriteString(rawData, p.ServAddr)
	binary.Write(rawData, binary.BigEndian, p.ServPort)
	ctypes.AsVarInt(p.NextState).WriteVarInt(rawData)

	return rawData.Bytes()
}

func (p *HandshakePacket) FillFromRaw(raw []byte) error {
	r := bytes.NewReader(raw)

	protoVersion, err := ctypes.ReadVarInt(r)

	if err == io.EOF {
		p.StatusReq = true
		return nil
	}

	servAddr, err := ctypes.ReadString(r)
	var servPortRaw [2]byte
	_, err = r.Read(servPortRaw[:])
	servPort := binary.BigEndian.Uint16(servPortRaw[:])
	nextState, err := ctypes.ReadVarInt(r)

	if err != nil {
		return err
	}

	p.ProtoVersion = protoVersion.Value()
	p.ServAddr = servAddr
	p.ServPort = servPort
	p.NextState = nextState.Value()

	return nil
}
