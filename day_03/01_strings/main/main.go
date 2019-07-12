package main

import (
	"fmt"
	"strings"
)

func main() {
	one := "Hello World, Hello Golang"

	// 测试字符串是否以指定字符开头 返回bool
	fmt.Println(strings.HasPrefix(one, "Hello"))

	// 测试字符串是否以指定字符结尾 返回bool
	fmt.Println(strings.HasSuffix(one, "World"))

	// 判断指定字符在字符串中首次出现的位置, 如果没有出现, 返回-1
	fmt.Println(strings.Index(one, "Go"))

	// 判断指定字符在字符串中最后出现的位置, 如果没有出现, 返回-1
	fmt.Println(strings.LastIndex(one, "Hello"))

	// 字符串替换, 最后一位参数是替换几次(必填, 替换所有填0或-1)
	one = strings.Replace(one, "Hello", "你好", 1)
	fmt.Println(one)
}
