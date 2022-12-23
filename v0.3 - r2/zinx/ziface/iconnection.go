package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	Send(data []byte)
	GetRemoteAddr() string
	GetTcpConn() *net.TCPConn
}