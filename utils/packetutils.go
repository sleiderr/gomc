package utils

type FieldlessPacket struct{}

func (p *FieldlessPacket) Raw() []byte {
	return make([]byte, 0)
}

func (p *FieldlessPacket) FillFromRaw(raw []byte) error {
	return nil
}
