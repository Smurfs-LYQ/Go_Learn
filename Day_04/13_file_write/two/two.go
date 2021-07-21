package two

import (
	"bufio"
	"fmt"
	"os"
)

// Two bufio.NewWriter写入文件实例
func Two() {
	// 打开文件并设置打开权限
	file, err := os.OpenFile("../xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	// 关闭文件
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件失败，错误信息: ", err)
		return
	}

	// 实例化一个bufio对象
	writer := bufio.NewWriter(file)
	for i := 0; i <= 10; i++ {
		writer.WriteString(fmt.Sprintf("hello %d\n", i)) // 将数据先写入到缓存中
	}
	writer.Flush() // 将缓存中的内容写入文件
}
