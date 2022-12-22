package main

import (
	"zinx_practice/zinx/znet"
)

func main() {
	server := &znet.Server{
		Name: "Game Server",
		IP:   "0.0.0.0",
		Port: 8888,
	}

	server.Serve()
}
