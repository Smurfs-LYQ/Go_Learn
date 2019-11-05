package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开第一个文件，确保文件是在的
	file_1, err := os.Open("./a.txt")
	defer file_1.Close()
	if err != nil {
		fmt.Println("打开文件失败，错误信息: ", err)
		return
	}

	// 打开第二个文件，确保文件存在，如果不存在，创建文件
	file_2, err := os.OpenFile("./b.txt", os.O_CREATE|os.O_WRONLY, 0644)
	defer file_2.Close()
	if err != nil {
		fmt.Println("打开文件失败，错误信息: ", err)
		return
	}

	num, err := io.Copy(file_2, file_1)
	if err != nil {
		fmt.Println("文件拷贝失败，失败原因: ", err)
		return
	}
	fmt.Println("文件拷贝成功，共拷贝字节数", num)
}
