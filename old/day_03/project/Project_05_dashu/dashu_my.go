package Project_05_dashu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func dashu(num []string) (result string) {
	// 去除字符串中的空格
	num1 := strings.TrimSpace(num[0])
	num2 := strings.TrimSpace(num[1])

	// 获取两个字符串的长度
	index1 := len(num1) - 1
	index2 := len(num2) - 1
	left := 0

	for index1 >= 0 && index2 >= 0 {
		c1 := num1[index1] - '0'
		c2 := num2[index2] - '0'

		sum := int(c1+c2) + left

		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}

		result = fmt.Sprintf("%d%s", sum%10, result)

		index1--
		index2--
	}

	for index1 >= 0 {
		c1 := num1[index1] - '0'

		sum := int(c1) + left

		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}

		result = fmt.Sprintf("%d%s", sum%10, result)

		index1--
	}

	for index2 >= 0 {
		c2 := num2[index2] - '0'

		sum := int(c2) + left

		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}

		result = fmt.Sprintf("%d%s", sum%10, result)

		index2--
	}

	if left == 1 {
		result = fmt.Sprintf("1%s", result)
	}

	return
}

// 计算两个大数相加的和，这两个大数会超过int64的表示范围
func Project_dashu_1() {
	// 从终端读取一行的内容
	reader := bufio.NewReader(os.Stdin) // bufio是带缓冲区的IO, NewReader是初始化一个读的实例, os.Stdin是从标准的终端输入读取
	result, _, err := reader.ReadLine() // 这个是在读的实例中调用单独读取一行的方法

	if err != nil {
		fmt.Println("read from console err: ", err)
		return
	}

	fmt.Println(dashu(strings.Split(string(result), "+")))
}
