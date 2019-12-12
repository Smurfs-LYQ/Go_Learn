package main

import "net"

import "fmt"

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		// IP: "127.0.0.1",
		IP: net.IPv4(0, 0, 0, 0), // 监听本机
		// IP:   net.ParseIP("127.0.0.1"), // 效果同上 ParseIP可以把字符串类型的IP转换为IP类型
		Port: 30000,
	})

	if err != nil {
		fmt.Println("启动server失败，err: ", err)
		return
	}

	defer listener.Close()

	for {
		var buf [1024]byte
		n, addr, err := listener.ReadFromUDP(buf[:]) // 参数: 用于接收客户端发送数据的切片 返回值: 1. 客户端发送数据的字节数, 2. 客户端标识(因为UDP的特性导致谁都可以给服务器端发送数据，addr返回值就可以辨识是谁发送的数据，也需要通过它为客户端返回信息) 3. 错误信息
		if err != nil {
			fmt.Println("接收消息失败, err: ", err)
			continue
		}

		// 打印客户端发送来的消息
		fmt.Printf("接收到来自%v的消息: %v\n", addr, string(buf[:n]))

		// 为客户端回复信息
		_, err = listener.WriteToUDP([]byte("收到了"), addr) // 参数: 1. 回复的信息内容  2. 客户端标识   返回值: 1. 发送过去的字节数 2. 错误信息
		if err != nil {
			fmt.Printf("为客户端%v发送消息失败，err: %v\n", addr, err)
			continue
		}
	}
}
