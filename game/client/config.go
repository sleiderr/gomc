package client

import (
	"fmt"

	"github.com/sleiderr/gomc/cnet/packet"
	"github.com/sleiderr/gomc/cnet/packet/config"
	"github.com/sleiderr/gomc/utils"
)

func (c *Client) handleConfiguration(rPacket *packet.CraftPacket) error {
	if rPacket.Id().Id == packet.ClientConfigPacketId {
		cPacket, ok := rPacket.Payload().(*config.ClientInformation)

		if !ok {
			return fmt.Errorf("Invalid client configuration packet")
		}

		_ = cPacket.ViewDistance

		finishPacket := packet.NewCraftPacket(packet.NewPacketType(byte(c.state), packet.FinishConfigPacketId), &utils.FieldlessPacket{})

		c.raw.Write(finishPacket.AsRaw().Bytes())
	}

	if rPacket.Id().Id == packet.FinishConfigPacketId {
		if c.login == nil {
			return fmt.Errorf("Unexpected packet received")
		}

		c.state = Play
		c.login.Status = LoggedIn

		fmt.Printf("%s just joined the game !\n", c.login.Username)
	}

	if rPacket.Id().Id == packet.ConfigPluginMessage {
		mPacket, ok := rPacket.Payload().(*config.PluginMessage)

		if !ok {
			return fmt.Errorf("Invalid plugin message received")
		}

		fmt.Println(mPacket.Identifier)
		fmt.Println(mPacket.Data)
	}

	return nil
}
