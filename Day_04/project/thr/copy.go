package thr

import (
	"fmt"
	"io"
	"os"
)

// Copy 借助IO.Copy实现一个拷贝文件的函数
func Copy(srcName, dstName string) {
	// 打开第一个文件，确保文件是在的
	file_1, err := os.Open(srcName)
	defer file_1.Close()
	if err != nil {
		fmt.Println("打开文件失败，错误信息: ...", err)
		return
	}

	// 打开第二个文件，确保文件存在，如果不存在，创建文件
	file_2, err := os.OpenFile(dstName, os.O_CREATE|os.O_WRONLY, 0755)
	defer file_2.Close()
	if err != nil {
		fmt.Println("打开文件失败，错误信息: ", err)
		return
	}

	num, err := io.Copy(file_2, file_1) // 参数1: 目标文件    参数2: 被拷贝的文件
	if err != nil {
		fmt.Println("文件拷贝失败，失败原因: ", err)
		return
	}
	fmt.Println("文件拷贝成功，共拷贝字节数", num)
}
