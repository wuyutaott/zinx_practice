package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp4", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("net.Listen error:", err)
	}
	defer listener.Close()

	fmt.Println("监听端口成功")

	if _, err := listener.Accept(); err != nil {
		fmt.Println("accept error:", err)
	}
	fmt.Println("收到一个连接")

	time.Sleep(3 * time.Second)
}
