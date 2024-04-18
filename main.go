package main

import "github.com/sleiderr/gomc/cnet"

func main() {
	server := cnet.NewServer("0.0.0.0", 25565)

	_ = server.ListenAndServe()

	for {
	}
}
