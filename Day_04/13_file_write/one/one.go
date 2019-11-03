package one

import (
	"fmt"
	"os"
)

func One() {
	// 打开文件支持文件写入
	file, err := os.OpenFile("../xx.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755) // 多个打开模式使用 | 来分隔
	// 关闭文件
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败，错误信息: ", err)
		return
	}
	file.Write([]byte("hello world\n"))
	file.WriteString("Hello Golang\n")
}
