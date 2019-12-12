package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

/*
服务端处理流程:
	1. 监听端口
	2. 接收客户端请求建立连接
	3. 创建goroutine处理链接
*/

func main() {
	// 1. 监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:20000") // 参数: 1. 网络类型(tcp/udp)  2. IP地址:端口号
	if err != nil {
		fmt.Println("监听 :20000 失败，err: ", err)
		return
	}
	defer listener.Close() // 程序结束关闭监听

	// 2. 接收客户端请求建立连接
	for {
		// 等待客户端连接, 如果没有客户端连接就阻塞, 一直在等待
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("连接失败, err: ", err)
			continue
		}

		fmt.Println("与客户端连接成功, 用户已上线")

		// 3. 创建goroutine处理链接
		go process(conn)
	}
}

// process 单独处理链接的函数
func process(conn net.Conn) {
	// 单独的goroutine结束之后关闭连接
	defer conn.Close()

	// 从连接中接收数据
	/*
		var buf [1024]byte
		n, err := conn.Read(buf[:]) // 跟读文件一样，参数: 用于存放读取内容的byte数组  返回值: 读了多少数据, 错误信息

		if err == io.EOF {
			fmt.Println("对方已关闭连接")
			return
		} else if err != nil {
			fmt.Println("接收客户端发来的消息失败, err: ", err)
			return
		}

		fmt.Println("接收客户端发来的消息: ", string(buf[:n]))
	*/

	reader := bufio.NewReader(conn)
	res, err := reader.ReadString('\n')

	if err == io.EOF {
		fmt.Println("对方已关闭连接")
		return
	} else if err != nil {
		fmt.Println("接收客户端发来的消息失败, err: ", err)
		return
	}

	fmt.Println("接收客户端发来的消息: ", res)
}
