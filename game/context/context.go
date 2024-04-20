package context

import (
	"github.com/sleiderr/gomc/cnet/packet"
	"github.com/sleiderr/gomc/game/client"
)

type Context struct {
	client *client.Client
	packet *packet.CraftPacket
}

func NewContext(client *client.Client, packet *packet.CraftPacket) *Context {
	return &Context{client, packet}
}

func (ctx *Context) Client() *client.Client {
	return ctx.client
}

func (ctx *Context) Packet() *packet.CraftPacket {
	return ctx.packet
}
