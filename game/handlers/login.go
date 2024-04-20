package handlers

import (
	"fmt"

	"github.com/sleiderr/gomc/cnet/packet"
	"github.com/sleiderr/gomc/cnet/packet/login"
	"github.com/sleiderr/gomc/game/client"
	"github.com/sleiderr/gomc/game/context"
)

func RegisterLoginHandlers() {
	RegisterHandler(packet.NewPacketType(client.Login, packet.LoginStartPacketId), handleLoginStart)
	RegisterHandler(packet.NewPacketType(client.Login, packet.LoginAckPacketId), handleLoginAck)
}

func handleLoginStart(ctx *context.Context) error {
	lPacket, ok := ctx.Packet().Payload().(*login.LoginStart)

	if !ok {
		return fmt.Errorf("Invalid Login packet received")
	}

	ctx.Client().Login = new(client.LoginTransaction)
	ctx.Client().Login.Username = lPacket.Name
	ctx.Client().Login.PlayerUUID = lPacket.PlayerUuid

	// TODO: move online mode config some place else
	if !client.OnlineMode {
		ctx.Client().Login.Status = client.LoginSuccess
		ctx.Client().DispatchLoginSuccess()
	}

	return nil
}

func handleLoginAck(ctx *context.Context) error {
	if ctx.Client().Login == nil || ctx.Client().Login.Status != client.LoginSuccess {
		return fmt.Errorf("Unexpected login acknowledgment")
	}

	ctx.Client().Login.Status = client.ConfigurationOngoing
	ctx.Client().SetState(client.Configuration)

	return nil
}
