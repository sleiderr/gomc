package client

type ClientState byte

const (
	Handshake = iota
	Status
	Login
	Configuration
	Play
	Closed
)
