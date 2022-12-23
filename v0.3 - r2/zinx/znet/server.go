package znet

import (
	"fmt"
	"net"
	"zinx_practice/zinx/ziface"
)

type Server struct {
	Name   string
	IP     string
	Port   int
	Router ziface.IRoute
}

func (s *Server) Start() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("ResolveTCPAddr err:", err)
		return
	}
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	fmt.Printf("server [%s] start %s:%d \n", s.Name, s.IP, s.Port)

	var cid uint32 = 0

	for {
		tcpConn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}

		conn := NewConnection(cid, tcpConn, s.Router)
		go conn.Start()

		cid++
	}
}

func (s *Server) Stop() {
	fmt.Println("server stop")
}

func (s *Server) Serve() {
	go s.Start()
}

func (s *Server) SetRouter(router ziface.IRoute) {
	s.Router = router
}

func NewServer(name string, ip string, port int) ziface.IServer {
	return &Server{
		Name: name,
		IP:   ip,
		Port: port,
	}
}
