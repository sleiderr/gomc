package handlers

import (
	"github.com/sleiderr/gomc/cnet/packet"
	"github.com/sleiderr/gomc/game/context"
)

type PacketHandler func(*context.Context) error

var handlersRegistry = make(map[packet.CraftPacketType][]PacketHandler)

func InitDefaultHandlers() {
	RegisterConfigHandlers()
}

func RegisterHandler(pType packet.CraftPacketType, handler PacketHandler) {
	if handlers, ok := handlersRegistry[pType]; ok {
		handlersRegistry[pType] = append(handlers, handler)
	} else {
		handlersRegistry[pType] = make([]PacketHandler, 1)
		handlersRegistry[pType][0] = handler
	}
}

func GetHandlers(pType packet.CraftPacketType) []PacketHandler {
	var handlers []PacketHandler
	if sHandlers, ok := handlersRegistry[pType]; ok {
		handlers = sHandlers
	}

	return handlers
}
