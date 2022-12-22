package znet

import (
	"fmt"
	"net"
	"zinx_practice/zinx/ziface"
)

type MsgHandle func(conn ziface.IConnection, data []byte, len int)

type Connection struct {
	ID         int
	conn       net.Conn
	msgHandler MsgHandle
	closeChan  chan bool
}

// 开始从管道中读取数据
func (c *Connection) startRead() {
	for {
		buf := make([]byte, 512)
		n, err := c.conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		c.msgHandler(c, buf, n)
	}
}

// Start 启动连接
func (c *Connection) Start() {
	go c.startRead()
}

// Stop 停止连接
func (c *Connection) Stop() {
	// 关闭连接
	c.conn.Close()

	// 关闭管道
	close(c.closeChan)
}

// Send 发送消息
func (c *Connection) Send(data []byte) {
	c.conn.Write(data)
}

// GetRemoteAddr 获取对端地址
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

// NewConnection 新建一个连接
func NewConnection(ID int, conn net.Conn, msgHandler MsgHandle) ziface.IConnection {
	fmt.Println("新建连接 ID =", ID)
	return &Connection{
		ID:         ID,
		conn:       conn,
		msgHandler: msgHandler,
		closeChan:  make(chan bool),
	}
}
