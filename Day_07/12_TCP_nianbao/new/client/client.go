package main

import (
	"Go_Learn/Day_07/12_TCP_nianbao/new/proto"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("与服务器建立连接失败, err :", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := fmt.Sprintf("test%d", i)
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("信息编码失败, err :", err)
			return
		}
		conn.Write(data)
	}
}
