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
	ConfigurationOngoing
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

		c.Login = new(LoginTransaction)
		c.Login.Username = lPacket.Name
		c.Login.PlayerUUID = lPacket.PlayerUuid

		if !OnlineMode {
			c.Login.Status = LoginSuccess
			c.dispatchLoginSuccess()
		}
	}

	if rPacket.Id().Id == packet.LoginAckPacketId {
		if c.Login == nil || c.Login.Status != LoginSuccess {
			return fmt.Errorf("Unexpected login acknowledgment")
		}

		c.Login.Status = ConfigurationOngoing
		c.state = Configuration
	}

	return nil
}

func (c *Client) dispatchLoginSuccess() {
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
