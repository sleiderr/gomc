package cnet

import (
	"fmt"
	"io"
	"net"
	"strconv"

	"github.com/sleiderr/gomc/cnet/packet"
	"github.com/sleiderr/gomc/game/client"
	"github.com/sleiderr/gomc/game/context"
	"github.com/sleiderr/gomc/game/handlers"
)

type CraftServer struct {
	host string
	port int
}

func NewServer(host string, port int) *CraftServer {
	return &CraftServer{
		host: host,
		port: port,
	}
}

func (serv *CraftServer) ListenAndServe() error {
	addr, err := net.ResolveTCPAddr("tcp", serv.host+":"+strconv.Itoa(serv.port))

	if err != nil {
		return err
	}

	conn, err := net.ListenTCP("tcp", addr)

	if err != nil {
		return err
	}

	fmt.Printf("Started listening on %s:%d\n", serv.host, serv.port)

	go func() {
		defer conn.Close()
		packet.InitPacketTypes()
		handlers.InitDefaultHandlers()
		for {
			rclient, err := conn.AcceptTCP()

			if err != nil {
				continue
			}

			_ = rclient.SetKeepAlive(true)

			go handleClient(client.NewClient(rclient))
		}
	}()

	return nil
}

func handleClient(rClient *client.Client) {
	defer rClient.Conn().Close()
	fmt.Printf("Received incoming connection from %s\n", rClient.Conn().RemoteAddr())

	for {
		p, err := packet.ReadCraftPacket(rClient.Conn())

		if err == io.EOF {
			fmt.Println("Client disconnected")
			break
		}

		if err != nil {
			fmt.Println("Error while reading packet")
			break
		}

		cp, err := packet.ParseRaw(p, byte(rClient.State()))
		if err != nil {
			fmt.Printf("Error while reading packet: %s\n", err.Error())
			continue
		}

		rClient.HandlePacket(cp)
		ctx := context.NewContext(rClient, cp)
		for _, handler := range handlers.GetHandlers(cp.Id()) {
			handler(ctx)
		}
	}
}
