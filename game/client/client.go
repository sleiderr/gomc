package client

import (
	"net"

	"github.com/sleiderr/gomc/cnet/packet"
	"github.com/sleiderr/gomc/cnet/packet/slp"
)

type Client struct {
	raw   *net.TCPConn
	state ClientState
	Login *LoginTransaction
}

func NewClient(raw *net.TCPConn) *Client {
	return &Client{
		raw:   raw,
		state: Handshake,
	}
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

func (c *Client) SetState(state ClientState) {
	c.state = state
}

func (c *Client) Conn() *net.TCPConn {
	return c.raw
}
