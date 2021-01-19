package main

import (
	"bradenn/config"
	"bradenn/server"
)

func main() {
	config.Init()
	server.Serve()
}
