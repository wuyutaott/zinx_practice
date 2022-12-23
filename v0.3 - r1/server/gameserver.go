package main

import (
	"fmt"
	"zinx_practice/zinx/ziface"
	"zinx_practice/zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (r *PingRouter) PrevHandle(request ziface.IRequest) {
	fmt.Println("before ping")
	request.GetConn().Send([]byte("before ping"))
}

func (r *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("ping")
	request.GetConn().Send([]byte("ping"))
}

func (r *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("after ping")
	request.GetConn().Send([]byte("after ping"))
}

func main() {
	router := &PingRouter{}
	server := znet.NewServer("Game Server", "0.0.0.0", 8888, router)

	server.Serve()
}
