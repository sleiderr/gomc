package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/sleiderr/gomc/cnet/packet"
	"github.com/sleiderr/gomc/cnet/packet/slp"
	"github.com/sleiderr/gomc/game"
	"github.com/sleiderr/gomc/game/client"
	"github.com/sleiderr/gomc/game/context"
)

func RegisterStatusHandlers() {
	RegisterHandler(packet.NewPacketType(client.Handshake, packet.HandshakePacketId), handleHandshake)
	RegisterHandler(packet.NewPacketType(client.Status, packet.HandshakePacketId), handleHandshake)
	RegisterHandler(packet.NewPacketType(client.Status, packet.PingPacketId), handlePing)
}

func handlePing(ctx *context.Context) error {
	if pPacket, ok := ctx.Packet().Payload().(*slp.PingPacket); ok {
		ctx.Client().Pong(pPacket)
	}

	return nil
}

func handleHandshake(ctx *context.Context) error {
	hPacket, ok := ctx.Packet().Payload().(*packet.HandshakePacket)

	if !ok {
		return fmt.Errorf("Received unexpected packet during handshake")
	}

	if hPacket.StatusReq {
		status, err := json.Marshal(game.GetGameStatus())

		if err != nil {
			return err
		}

		statusResp := packet.NewCraftPacket(packet.NewPacketType(byte(ctx.Client().State()), packet.HandshakePacketId), &packet.StatusResponsePacket{Status: string(status)})
		ctx.Client().Conn().Write(statusResp.AsRaw().Bytes())
	} else {
		ctx.Client().SetState(client.ClientState(hPacket.NextState))
	}

	return nil
}
