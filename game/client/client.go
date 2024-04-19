package client

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/sleiderr/gomc/cnet/packet"
	"github.com/sleiderr/gomc/game"
)

type Client struct {
	raw   *net.TCPConn
	state ClientState
}

func NewClient(raw *net.TCPConn) *Client {
	return &Client{
		raw:   raw,
		state: Handshake,
	}
}

func (c *Client) HandlePacket(rPacket *packet.CraftPacket) error {
	var err error
	switch c.State() {
	case Handshake:
		err = c.handleHandshake(rPacket)
	case Status:
		if rPacket.Id() == packet.Handshake {
			err = c.handleHandshake(rPacket)
		}
	case Login:
		if rPacket.Id() == packet.Handshake {
			err = c.handleHandshake(rPacket)
		}
	}

	return err
}

func (c *Client) handleHandshake(rPacket *packet.CraftPacket) error {
	hPacket, ok := rPacket.Payload().(*packet.HandshakePacket)

	if rPacket.Id() != packet.Handshake || !ok {
		return fmt.Errorf("Received unexpected packet during handshake")
	}

	if hPacket.StatusReq {
		status, err := json.Marshal(game.GetGameStatus())

		if err != nil {
			return err
		}

		statusResp := packet.NewCraftPacket(packet.Handshake, &packet.StatusResponsePacket{Status: string(status)})
		c.raw.Write(statusResp.AsRaw().Bytes())
	}

	return nil
}

func (c *Client) NextState() {
	if c.state == Closed {
		return
	}
	c.state += 1
}

func (c *Client) State() ClientState {
	return c.state
}

func (c *Client) Conn() *net.TCPConn {
	return c.raw
}
