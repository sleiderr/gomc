package slp

import (
	"bytes"
	"encoding/binary"
)

type PingPacket struct {
	Payload uint64
}

func (p *PingPacket) Raw() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, p.Payload)

	return buf.Bytes()
}

func (p *PingPacket) FillFromRaw(raw []byte) error {
	p.Payload = binary.BigEndian.Uint64(raw)
	return nil
}
