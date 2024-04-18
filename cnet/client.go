package cnet

import "net"

type Client struct {
	raw *net.TCPConn
}

func NewClient(raw *net.TCPConn) Client {
	return Client{
		raw: raw,
	}
}
