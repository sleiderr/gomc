package client

import (
	"github.com/google/uuid"
	"github.com/sleiderr/gomc/cnet/packet"
	"github.com/sleiderr/gomc/cnet/packet/login"
)

type LoginStatus byte

var OnlineMode bool = false

const (
	LoginStart LoginStatus = iota
	AwaitingEncryptionResponse
	LoginSuccess
	ConfigurationOngoing
	LoggedIn
)

type LoginTransaction struct {
	Status     LoginStatus
	Username   string
	PlayerUUID uuid.UUID
}

func (c *Client) DispatchLoginSuccess() {
	if c.Login == nil || c.Login.Status != LoginSuccess {
		return
	}

	loginResp := packet.NewCraftPacket(packet.NewPacketType(byte(c.state), packet.LoginSuccessPacketId), &login.LoginSuccess{
		PlayerUUID:      c.Login.PlayerUUID,
		Username:        c.Login.Username,
		PropertiesCount: 0,
		Properties:      make([]login.LoginProperty, 0),
	})

	c.raw.Write(loginResp.AsRaw().Bytes())
}
