package main

import (
	"fmt"
)

// Reverse 字符串翻转
func Reverse(s string) {
	var res string
	for i := len(s); i > 0; i-- {
		res = res + s[i-1:i]
		// 注意:使用切片读取字符串中的单个字符的时候返回值为那个字符对应的Unicode码
		// res = res + fmt.Sprintf("%c", s[i-1])
	}
	fmt.Println(res)
}

func main() {
	// 字符串拼接
	var (
		str1 = "hello"
		str2 = "golang"
	)

	// 字符串拼接
	res1 := str1 + " " + str2
	fmt.Println(res1)

	// 字符串占位符拼接
	res2 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(res2)

	// 查看字符串长度
	fmt.Println(len(res2))

	// 字符串切片
	fmt.Println(res1[:5])
	fmt.Println(res1[6:])

	// 字符串翻转
	Reverse(res1)
}
