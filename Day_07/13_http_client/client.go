package main

import (
	"fmt"
	"io"
	"net"
)

// HTTP Client

func main() {
	// 建立拨号网络
	conn, err := net.Dial("tcp", "www.liwenzhou.com:443")
	if err != nil {
		fmt.Println("拨号网络建立失败, err: ", err)
	}
	defer conn.Close()

	// 发送请求
	/*
		GET / HTTP/1.1\r\n\r\n
		GET 	 : 请求类型
		/ 		 : 请求路径
		HTTP/1.1 : 协议/协议版本
	*/
	// fmt.Fprintf(conn, "GET / HTTP/1.1\r\n\r\n")
	// 向连接中写入数据
	conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	// 接收数据
	var buf [1014]byte
	for {
		// 从连接中读取数据
		n, err := conn.Read(buf[:])

		// 判断文件是否读完
		if err == io.EOF {
			return
		}
		// 判断文件返回信息
		if err != nil {
			fmt.Println("获取返回信息失败, err: ", err)
			return
		}

		fmt.Println(string(buf[:n]))
	}
}
