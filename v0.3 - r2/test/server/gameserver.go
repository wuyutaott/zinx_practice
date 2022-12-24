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

func (pr *PingRouter) Handle(req ziface.IRequest) {
	msg := req.GetMsg()
	fmt.Printf("收到客户端消息 id = %d, data = %s \n", msg.GetID(), string(msg.GetData()))

	msg2 := &znet.Message{
		ID: 2,
		DataLen: 5,
		Data: []byte("world"),
	}
	if err := req.GetConn().Send(msg2); err != nil {
		fmt.Println("send msg err:", err)
	}

	msg3 := &znet.Message{
		ID: 3,
		DataLen: 5,
		Data: []byte("12345"),
	}
	if err := req.GetConn().Send(msg3); err != nil {
		fmt.Println("send msg err:", err)
	}
}

func main()  {
	server := znet.NewServer("mmo_server", "0.0.0.0", 8888)
	server.SetRouter(&PingRouter{})
	server.Serve()
	for {
		time.Sleep(1 * time.Second)
	}
}