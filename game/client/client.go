package client

import (
	"math/rand/v2"
	"net"

	"github.com/sleiderr/gomc/cnet/packet"
	"github.com/sleiderr/gomc/cnet/packet/gamepacket"
	"github.com/sleiderr/gomc/cnet/packet/slp"
	"github.com/sleiderr/gomc/game"
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

func (c *Client) SetSpawnPosition(newPos game.EntityLocation, angle float32) {
	spawnPosPacket := packet.NewCraftPacket(packet.NewPacketType(byte(c.State()), packet.PlaySetSpawnPos), &gamepacket.DefaultSpawnPosition{
		Location: newPos,
		Angle:    angle,
	})

	c.Conn().Write(spawnPosPacket.AsRaw().Bytes())
}

func (c *Client) Teleport(newPos game.EntityPosition, rel bool) {
	var flags byte
	if rel {
		flags = 0b11111
	}

	tpPacket := packet.NewCraftPacket(packet.NewPacketType(byte(c.State()), packet.PlaySyncPosition), &gamepacket.SynchronizePositionPacket{
		Position:   newPos,
		Flags:      gamepacket.TeleportFlags(flags),
		TeleportID: rand.Int32(),
	})

	c.Conn().Write(tpPacket.AsRaw().Bytes())
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
