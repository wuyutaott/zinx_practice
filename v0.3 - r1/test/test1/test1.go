package main

import (
	"fmt"
	"net"
	"time"
)

func RunClient() {
	conn, err := net.Dial("tcp4", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	for {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			fmt.Println("client write err:", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func DelayCloseServer(conn net.Conn) {
	time.Sleep(5 * time.Second)
	fmt.Println("服务器主动断开与客户端的连接")
	conn.Close()
}

func RunServer() {
	listener, err := net.Listen("tcp4", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			return
		}

		go DelayCloseServer(conn)

		for {
			buf := make([]byte, 512)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("read err:", err)
				break
			}
			fmt.Println("收到数据：", string(buf[:n]))
		}
	}()
}

func main() {
	go RunServer()

	go RunClient()

	exitChan := make(chan bool, 1)

	go func() {
		// 等待一段时间后退出程序
		time.Sleep(time.Second * 10)
		exitChan <- true
	}()

	<-exitChan
}
