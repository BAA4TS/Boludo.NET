package main

import (
	"BOLUDO.NET/nets"
	"BOLUDO.NET/repl"
)

func main() {
	go nets.Start()
	repl.Repl()
}
