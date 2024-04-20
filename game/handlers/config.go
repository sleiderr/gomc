package handlers

import (
	"fmt"

	"github.com/sleiderr/gomc/cnet/packet"
	"github.com/sleiderr/gomc/cnet/packet/config"
	"github.com/sleiderr/gomc/game/client"
	"github.com/sleiderr/gomc/game/context"
	"github.com/sleiderr/gomc/utils"
)

func RegisterConfigHandlers() {
	RegisterHandler(packet.NewPacketType(client.Configuration, packet.ConfigPluginMessage), handlePluginMessage)
	RegisterHandler(packet.NewPacketType(client.Configuration, packet.ClientConfigPacketId), handleClientConfig)
	RegisterHandler(packet.NewPacketType(client.Configuration, packet.FinishConfigPacketId), handleFinishConfig)
}

func handleClientConfig(ctx *context.Context) error {
	cPacket, ok := ctx.Packet().Payload().(*config.ClientInformation)

	if !ok {
		return fmt.Errorf("Invalid client configuration packet")
	}

	_ = cPacket.ViewDistance

	finishPacket := packet.NewCraftPacket(packet.NewPacketType(byte(ctx.Client().State()), packet.FinishConfigPacketId), &utils.FieldlessPacket{})

	ctx.Client().Conn().Write(finishPacket.AsRaw().Bytes())

	return nil
}

func handleFinishConfig(ctx *context.Context) error {
	if ctx.Client().Login == nil {
		return fmt.Errorf("Unexpected packet received")
	}

	ctx.Client().SetState(client.Play)
	ctx.Client().Login.Status = client.LoggedIn

	fmt.Printf("%s just joined the game !\n", ctx.Client().Login.Username)

	return nil
}

func handlePluginMessage(ctx *context.Context) error {
	mPacket, ok := ctx.Packet().Payload().(*config.PluginMessage)

	if !ok {
		return fmt.Errorf("Invalid plugin message received")
	}

	fmt.Println(mPacket.Identifier)
	fmt.Println(mPacket.Data)

	return nil
}
