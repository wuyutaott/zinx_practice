package znet

import (
	"fmt"
	"io"
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
		header := make([]byte, 8)
		if _, err := io.ReadFull(c.tcpConn, header); err != nil {
			c.closeChan <- true
			fmt.Println("read header err:", err)
			return
		}

		dp := NewDataPack()
		msg, err := dp.UnPack(header)
		if err != nil {
			fmt.Println("UnPack err:", err)
			c.closeChan <- true
			return
		}

		body := make([]byte, msg.GetDataLength())
		if _, err := io.ReadFull(c.tcpConn, body); err != nil {
			c.closeChan <- true
			fmt.Println("read body err:", err)
			return
		}
		msg.SetData(body)

		req := NewRequest(c, msg)

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

func (c *Connection) Send(msg ziface.IMessage) error {
	dp := NewDataPack()
	data, err := dp.Pack(msg)
	if err != nil {
		return err
	}
	if _, err := c.tcpConn.Write(data); err != nil {
		return err
	}
	return nil
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
