package client

import (
	"fmt"

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
	LoggedIn
)

type LoginTransaction struct {
	Status     LoginStatus
	Username   string
	PlayerUUID uuid.UUID
}

func (c *Client) handleLogin(rPacket *packet.CraftPacket) error {
	if rPacket.Id().Id == packet.LoginStartPacketId {
		lPacket, ok := rPacket.Payload().(*login.LoginStart)

		if !ok {
			return fmt.Errorf("Invalid Login packet received")
		}

		c.login = new(LoginTransaction)
		c.login.Username = lPacket.Name
		c.login.PlayerUUID = lPacket.PlayerUuid

		if !OnlineMode {
			c.login.Status = LoginSuccess
			c.dispatchLoginSuccess()
		}
	}

	if rPacket.Id().Id == packet.LoginAckPacketId {
		if c.login == nil || c.login.Status != LoginSuccess {
			return fmt.Errorf("Unexpected login acknowledgment")
		}

		c.login.Status = LoggedIn
		c.state = Configuration
	}

	return nil
}

func (c *Client) dispatchLoginSuccess() {
	if c.login == nil || c.login.Status != LoginSuccess {
		return
	}

	loginResp := packet.NewCraftPacket(packet.NewPacketType(byte(c.state), packet.LoginSuccessPacketId), &login.LoginSuccess{
		PlayerUUID:      c.login.PlayerUUID,
		Username:        c.login.Username,
		PropertiesCount: 0,
		Properties:      make([]login.LoginProperty, 0),
	})

	c.raw.Write(loginResp.AsRaw().Bytes())
}
