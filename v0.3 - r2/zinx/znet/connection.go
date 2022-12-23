package znet

import (
	"fmt"
	"net"
	"zinx_practice/zinx/ziface"
)

type Connection struct {
	id        uint32
	tcpConn   *net.TCPConn
	closeChan chan bool
	router    ziface.IRoute
}

func (c *Connection) startRead() {
	for {
		buf := make([]byte, 512)
		_, err := c.tcpConn.Read(buf)
		if err != nil {
			c.closeChan <- true
			return
		}

		req := NewRequest(c, buf)

		go func() {
			if c.router != nil {
				c.router.PrevHandle(req)
				c.router.Handle(req)
				c.router.PostHandle(req)
			}
		}()
	}
}

func (c *Connection) Start() {
	fmt.Println("connection start id:", c.id)
	defer c.Stop()
	go c.startRead()
	<-c.closeChan
}

func (c *Connection) Stop() {
	fmt.Println("connection stop id:", c.id)
	c.tcpConn.Close()
	close(c.closeChan)
}

func (c *Connection) Send(data []byte) {
	_, err := c.tcpConn.Write(data)
	if err != nil {
		fmt.Println("send err:", err)
	}
}

func (c *Connection) GetRemoteAddr() string {
	return c.tcpConn.RemoteAddr().String()
}

func (c *Connection) GetTcpConn() *net.TCPConn {
	return c.tcpConn
}

func NewConnection(id uint32, tcpConn *net.TCPConn, router ziface.IRoute) ziface.IConnection {
	return &Connection{
		id:        id,
		tcpConn:   tcpConn,
		closeChan: make(chan bool, 1),
		router:    router,
	}
}
