package znet

import (
	"fmt"
	"net"
	"zinx_practice/zinx/ziface"
)

type Server struct {
	// 服务器的名称
	Name string
	// 服务器绑定的IP地址
	IP string
	// 服务器的端口号
	Port int
	// 消息路由
	Router ziface.IRouter
}

func (s *Server) Start() {
	fmt.Printf("服务器启动，IP: %s, Port: %d\n", s.IP, s.Port)
	listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}

	fmt.Println("服务器启动监听")
	go func() {
		ID := 0
		for {
			tcpConn, err := listener.Accept()
			if err != nil {
				fmt.Println("accept err:", err)
				continue
			}

			fmt.Println("收到客户端连接")

			connection := NewConnection(ID, tcpConn, s.Router)
			go connection.Start()

			ID++
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	for {
		select {}
	}
}

func NewServer(name string, ip string, port int, router ziface.IRouter) ziface.IServer {
	return &Server{
		Name: name,
		IP: ip,
		Port: port,
		Router: router,
	}
}
