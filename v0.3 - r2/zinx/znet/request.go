package znet

import "zinx_practice/zinx/ziface"

type Request struct {
	conn ziface.IConnection
	msg ziface.IMessage
}

func (r *Request) GetConn() ziface.IConnection {
	return r.conn
}

func (r *Request) GetMsg() ziface.IMessage {
	return r.msg
}

func NewRequest(conn ziface.IConnection, msg ziface.IMessage) ziface.IRequest {
	return &Request{
		conn: conn,
		msg: msg,
	}
}
