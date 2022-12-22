package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("我是客户端")

	conn, err := net.Dial("tcp4", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}

	conn.Write([]byte{'h', 'e', 'l', 'l', 'o'})

	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("收到服务器消息：", buf[:n])

		time.Sleep(1 * time.Second)

		conn.Write([]byte{'1', '2', '3'})
	}
}
