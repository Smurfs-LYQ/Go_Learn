package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/*
客户端处理流程
	1. 建立与服务器的链接
	2. 进行数据收发
	3. 关闭链接
*/

func main() {
	for {
		// 1. 建立与服务器的链接
		conn, err := net.Dial("tcp", "127.0.0.1:20000") // 参数: 1. 网络类型(tcp/udp)  2. IP地址:端口号
		if err != nil {
			fmt.Println("连接服务端失败, err: ", err)
			return
		}

		// 关闭链接
		defer conn.Close()

		// 2. 进行数据收发
		// val := "test"

		reader := bufio.NewReader(os.Stdin) // 从标准输入获取输入的内容
		val, err := reader.ReadString('\n') // 读取内容直到遇到 \n
		if err != nil {
			fmt.Println("获取用户输入信息失败")
			continue
		}

		_, err = fmt.Fprint(conn, val)
		// _, err = conn.Write([]byte(val)) // 参数: 要发送的内容  返回值: 1. 成功写入的字节数量 2. 错误信息
		if err != nil {
			fmt.Println("发送消息失败, err: ", err)
			continue
		}
	}
}
