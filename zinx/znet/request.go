package znet

import "zinx_practice/zinx/ziface"

type Request struct {
	Conn ziface.IConnection
	Data []byte
}

func (r *Request) GetConn() ziface.IConnection {
	return r.Conn
}

func (r *Request) GetData() []byte {
	return r.Data
}

func NewRequest(conn ziface.IConnection, data []byte) ziface.IRequest {
	return &Request{
		Conn: conn,
		Data: data,
	}
}
