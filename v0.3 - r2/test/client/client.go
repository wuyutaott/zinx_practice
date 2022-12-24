package main

import (
	"fmt"
	"net"
	"time"
	"zinx_practice/zinx/znet"
)

func main() {
	conn, err := net.Dial("tcp4", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}

	go func() {
		for {
			header := make([]byte, 8)
			if _, err := conn.Read(header); err != nil {
				fmt.Println("read header err:", err)
				return
			}
			dp := znet.NewDataPack()
			msg, err := dp.UnPack(header)
			if err != nil {
				fmt.Println("unpack err:", err)
				return
			}
			body := make([]byte, msg.GetDataLength())
			if _, err := conn.Read(body); err != nil {
				fmt.Println("read body err:", err)
				return
			}
			msg.SetData(body)
			fmt.Printf("收到服务器消息 id = %d, data = %s \n", msg.GetID(), string(msg.GetData()))
		}
	}()

	for {
		msg := &znet.Message{
			DataLen: 5,
			ID: 1,
			Data: []byte("hello"),
		}
		dp := znet.NewDataPack()
		data, err := dp.Pack(msg)
		if err != nil {
			fmt.Println("pack msg err:", err)
			return
		}
		if _, err := conn.Write(data); err != nil {
			fmt.Println("write err:", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}