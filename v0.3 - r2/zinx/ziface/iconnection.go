package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	Send(msg IMessage) error
	GetRemoteAddr() string
	GetTcpConn() *net.TCPConn
}