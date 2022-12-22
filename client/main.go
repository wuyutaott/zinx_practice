package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("我是客户端")

	_, err := net.Dial("tcp4", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}

}
