package main

import (
	"fmt"
	"zinx_practice/zinx/ziface"
	"zinx_practice/zinx/znet"
)

func msgProcess(conn ziface.IConnection, data []byte, len int) {
	fmt.Println("消息处理")
	conn.Send(data[0:len])
}

func main() {
	server := &znet.Server{
		Name: "Game Server",
		IP:   "0.0.0.0",
		Port: 8888,
		MsgHandler: msgProcess,
	}

	server.Serve()
}
