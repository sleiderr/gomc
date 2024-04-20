package handlers

import (
	"github.com/sleiderr/gomc/cnet/packet"
	"github.com/sleiderr/gomc/game/context"
)

type PacketHandler func(*context.Context) error

type packetRoute struct {
	StateBound bool
	PacketID   uint32
	Handler    PacketHandler
}

var handlersRegistry = make(map[byte][]packetRoute)

func InitDefaultHandlers() {
	RegisterStatusHandlers()
	RegisterLoginHandlers()
	RegisterConfigHandlers()
}

func RegisterStateHandler(pState byte, handler PacketHandler) {
	route := packetRoute{true, 0, handler}
	if handlers, ok := handlersRegistry[pState]; ok {
		handlersRegistry[pState] = append(handlers, route)
	} else {
		handlersRegistry[pState] = make([]packetRoute, 1)
		handlersRegistry[pState][0] = route
	}
}

func RegisterHandler(pType packet.CraftPacketType, handler PacketHandler) {
	route := packetRoute{false, pType.Id, handler}
	if handlers, ok := handlersRegistry[pType.State]; ok {
		handlersRegistry[pType.State] = append(handlers, route)
	} else {
		handlersRegistry[pType.State] = make([]packetRoute, 1)
		handlersRegistry[pType.State][0] = route
	}
}

func GetHandlers(pType packet.CraftPacketType) []packetRoute {
	var handlers []packetRoute
	if sHandlers, ok := handlersRegistry[pType.State]; ok {
		handlers = sHandlers
	}

	return handlers
}
