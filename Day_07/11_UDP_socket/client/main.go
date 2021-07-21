package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	for {
		conn, err := net.Dial("udp", "127.0.0.1:30000")
		if err != nil {
			fmt.Println("链接server失败, err: ", err)
			return
		}

		defer conn.Close()

		// 从终端获取用户标准输入
		reader := bufio.NewReader(os.Stdin)
		val, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("获取用户输入失败, err: ", err)
			continue
		}

		// 向server端发送信息
		_, err = conn.Write([]byte(val))
		if err != nil {
			fmt.Println("发送信息失败, err: ", err)
			continue
		}

		// 接收server端发送回来的信息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("接收服务器端发送的信息失败, err: ", err)
			continue
		}
		fmt.Println("收到回复: ", string(buf[:n]))
	}
}
