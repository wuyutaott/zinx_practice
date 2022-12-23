package znet

import (
	"fmt"
	"net"
	"zinx_practice/zinx/ziface"
)

type Connection struct {
	ID         int
	Conn       net.Conn
	Router 	   ziface.IRouter
	ExitChan   chan bool
}

// 开始从管道中读取数据
func (c *Connection) startRead() {
	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			c.ExitChan <- true
			return
		}

		req := NewRequest(c, buf)

		go func() {
			c.Router.PrevHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}()
	}
}

// Start 启动连接
func (c *Connection) Start() {
	go c.startRead()
	defer c.Stop()

	<-c.ExitChan
}

// Stop 停止连接
func (c *Connection) Stop() {
	// 关闭连接
	c.Conn.Close()

	// 关闭管道
	close(c.ExitChan)

	fmt.Println("连接关闭 ID =", c.ID)
}

// Send 发送消息
func (c *Connection) Send(data []byte) {
	c.Conn.Write(data)
}

// GetRemoteAddr 获取对端地址
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// NewConnection 新建一个连接
func NewConnection(ID int, conn net.Conn, router ziface.IRouter) ziface.IConnection {
	fmt.Println("连接建立 ID =", ID)
	return &Connection{
		ID:         ID,
		Conn:       conn,
		Router: 	router,
		ExitChan:   make(chan bool),
	}
}
