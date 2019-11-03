package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("./xx.txt")
	// 关闭文件
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件失败, 错误信息: ", err)
		return
	}

	// 读取文件
	/*
		var res [50]byte
		num, err := file.Read(res[:])
		if err == io.EOF { // 判断是否读到了文件的末尾
			fmt.Println("文件已经读完了")
			return
		}
		if err != nil {
			fmt.Println("打开文件失败, 错误信息: ", err)
			return
		}
		fmt.Printf("字节数: %d\n", num)
		fmt.Println(string(res[:]))
	*/

	// 循环读取文件
	var res [50]byte
	for {
		num, err := file.Read(res[:])
		if err == io.EOF { // 判断是否读到了文件的末尾
			fmt.Println("文件已经读完了")
			return
		}
		if err != nil {
			fmt.Println("打开文件失败, 错误信息: ", err)
			return
		}
		fmt.Printf("字节数: %d\n", num)
		fmt.Println(string(res[:]))
	}
}
