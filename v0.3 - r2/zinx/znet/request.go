package znet

import "zinx_practice/zinx/ziface"

type Request struct {
	conn ziface.IConnection
	data []byte
}

func (r *Request) GetConn() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}

func NewRequest(conn ziface.IConnection, data []byte) ziface.IRequest {
	return &Request{
		conn: conn,
		data: data,
	}
}
