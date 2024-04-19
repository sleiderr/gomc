package client

type ClientState byte

const (
	Handshake = iota
	Status
	Login
	Play
	Closed
)
