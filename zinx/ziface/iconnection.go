package ziface

import "net"

type IConnection interface {
	// 启动连接
	Start()
	// 停止连接
	Stop()
	// 发送消息
	Send([]byte)
	// 获取对端地址
	GetRemoteAddr() net.Addr
}
