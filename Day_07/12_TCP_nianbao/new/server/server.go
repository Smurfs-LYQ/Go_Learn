package main

import (
	"Go_Learn/Day_07/12_TCP_nianbao/new/proto"
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("获取客户端发来的信息失败, err :", err)
			return
		}
		fmt.Println("收到client发来的数据: ", msg)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("监听失败, err :", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("接收客户端请求连接失败, err:", err)
			continue
		}
		go process(conn)
	}
}
