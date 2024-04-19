package client

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/sleiderr/gomc/cnet/packet"
	"github.com/sleiderr/gomc/cnet/packet/slp"
	"github.com/sleiderr/gomc/game"
)

type Client struct {
	raw   *net.TCPConn
	state ClientState
	login *LoginTransaction
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
		if rPacket.Id().Id == packet.HandshakePacketId {
			err = c.handleHandshake(rPacket)
		} else {
			err = c.handleStatus(rPacket)
		}
	case Login:
		err = c.handleLogin(rPacket)
	}

	return err
}

func (c *Client) handleStatus(rPacket *packet.CraftPacket) error {
	if pPacket, ok := rPacket.Payload().(*slp.PingPacket); rPacket.Id().Id == packet.PingPacketId && ok {
		c.Pong(pPacket)
	}

	return nil
}

func (c *Client) handleHandshake(rPacket *packet.CraftPacket) error {
	hPacket, ok := rPacket.Payload().(*packet.HandshakePacket)

	if rPacket.Id().Id != packet.HandshakePacketId || !ok {
		return fmt.Errorf("Received unexpected packet during handshake")
	}

	if hPacket.StatusReq {
		status, err := json.Marshal(game.GetGameStatus())

		if err != nil {
			return err
		}

		statusResp := packet.NewCraftPacket(packet.NewPacketType(byte(c.state), packet.HandshakePacketId), &packet.StatusResponsePacket{Status: string(status)})
		c.raw.Write(statusResp.AsRaw().Bytes())
	} else {
		c.state = ClientState(hPacket.NextState)
	}

	return nil
}

func (c *Client) Pong(in *slp.PingPacket) {
	c.raw.Write(packet.NewCraftPacket(packet.NewPacketType(byte(c.state), 1), in).AsRaw().Bytes())
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
