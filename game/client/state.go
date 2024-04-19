package client

type ClientState byte

const (
	Handshake = 1
	Status    = iota + 1
	Login
	Play
	Closed
)
