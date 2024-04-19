package login

import (
	"bytes"

	"github.com/google/uuid"
	"github.com/sleiderr/gomc/ctypes"
)

type LoginStart struct {
	Name       string
	PlayerUuid uuid.UUID
}

type LoginSuccess struct {
	PlayerUUID      uuid.UUID
	Username        string
	PropertiesCount int32
	Properties      []LoginProperty
}

type LoginProperty struct {
	Name      string
	Value     string
	IsSigned  bool
	Signature string
}

type LoginAck struct{}

func (p *LoginStart) Raw() []byte {
	rawData := new(bytes.Buffer)

	ctypes.WriteString(rawData, p.Name)
	rawData.Write(p.PlayerUuid[:])

	return rawData.Bytes()
}

func (p *LoginStart) FillFromRaw(raw []byte) error {
	r := bytes.NewReader(raw)

	name, err := ctypes.ReadString(r)
	pUuidBytes := make([]byte, 16)
	_, err = r.Read(pUuidBytes)
	pUuid, err := uuid.FromBytes(pUuidBytes)

	if err != nil {
		return err
	}

	p.Name = name
	p.PlayerUuid = pUuid

	return nil
}

func (p *LoginSuccess) Raw() []byte {
	rawData := new(bytes.Buffer)

	rawData.Write(p.PlayerUUID[:])
	ctypes.WriteString(rawData, p.Username)
	ctypes.AsVarInt(0).WriteVarInt(rawData)

	return rawData.Bytes()
}

func (p *LoginSuccess) FillFromRaw(raw []byte) error {
	return nil
}

func (p *LoginAck) Raw() []byte {
	return make([]byte, 0)
}

func (p *LoginAck) FillFromRaw(raw []byte) error {
	return nil
}
