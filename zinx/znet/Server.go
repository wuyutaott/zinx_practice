package znet

import (
	"fmt"
	"net"
)

type Server struct {
	// 服务器的名称
	Name string
	// 服务器绑定的IP地址
	IP string
	// 服务器的端口号
	Port int
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
		for {
			_, err := listener.Accept()
			if err != nil {
				fmt.Println("accept err:", err)
				continue
			}

			fmt.Println("收到客户端连接")
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