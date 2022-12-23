package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp4", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}

	go func() {
		for {
			buf := make([]byte, 512)
			cnt, err := conn.Read(buf)
			if err != nil {
				fmt.Println("read err:", err)
				return
			}
			fmt.Println("recive server msg:", string(buf[:cnt]))
		}
	}()

	for {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			fmt.Println("write err:", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}