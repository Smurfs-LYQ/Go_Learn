package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("../10_read_file/xx.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件失败, 错误信息: ", err)
		return
	}

	// 利用缓冲区从文件读数据
	reader := bufio.NewReader(file) // 实例化一个bufio的对象
	for {
		str, err := reader.ReadString('\n') // \n 是一个字节，代表的是读到哪里停止
		if err == io.EOF {
			fmt.Println(str) // 打印最后一行数据，因为已经读到了文件的末尾，而文件的末尾没有\n，所以这里需要使用println
			fmt.Println("文件读完了")
			return
		}
		if err != nil {
			fmt.Println("读取文件内容失败, 错误信息: ", err)
			return
		}
		fmt.Print(str) // 因为ReadString被设置为读到\n就结束，所以代表str字符串的结尾处有一个\n，如果使用println函数，那么会多出来一行
	}
}
