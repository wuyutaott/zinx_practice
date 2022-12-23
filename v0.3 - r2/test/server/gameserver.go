package main

import (
	"fmt"
	"time"
	"zinx_practice/zinx/ziface"
	"zinx_practice/zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) PrevHandle(req ziface.IRequest) {
	fmt.Println("PrevHandle")
}

func (pr *PingRouter) Handle(req ziface.IRequest) {
	fmt.Println("Handle")
	req.GetConn().Send([]byte("你好"))
}

func (pr *PingRouter) PostHandle(req ziface.IRequest) {
	fmt.Println("PostHandle")
}

func main()  {
	server := znet.NewServer("mmo_server", "0.0.0.0", 8888)
	server.SetRouter(&PingRouter{})
	server.Serve()
	for {
		time.Sleep(1 * time.Second)
	}
}