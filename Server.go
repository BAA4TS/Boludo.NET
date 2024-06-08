package main

import (
	"BOLUDO.NET/net"
	"BOLUDO.NET/repl"
)

func main() {
	go net.Start()
	repl.Repl()
}
