package cnet

import (
	"fmt"
	"net"
	"strconv"
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
		for {
			client, err := conn.AcceptTCP()

			if err != nil {
				continue
			}

			_ = client.SetKeepAlive(true)

			go handleClient(NewClient(client))
		}
	}()

	return nil
}

func handleClient(client Client) {
	fmt.Printf("Received new connection from %s\n", client.raw.RemoteAddr().String())
}
