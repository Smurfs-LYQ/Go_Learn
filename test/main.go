package main

import (
	"fmt"
	"strings"
)

func main() {
	one := "Hello World, Hello Golang"

	// 检测字符串是否以指定字符开头
	fmt.Println(strings.HasPrefix(one, "Hello"))

	// 检测字符串是否以指定字符结尾
	fmt.Println(strings.HasSuffix(one, "Golang"))

	// 字符首次出现的位置
	fmt.Println(strings.Index(one, "Golang"))

	// 字符最后一次出现的位置
	fmt.Println(strings.LastIndex(one, "Hello"))

	// 字符串替换
	two := "Hello World, Hello Golang"
	fmt.Println(strings.Replace(two, "Hello", "你好", 2))

	// 字符串计数
	fmt.Println(strings.Count(one, "Hello"))

	// 字符串重复指定次数
	fmt.Println(strings.Repeat("哈", 3))

	// 转为小写
	fmt.Println(strings.ToLower("HELLO"))

	// 转为大写
	fmt.Println(strings.ToUpper("hello"))

	// 去除字符串首尾的空白字符
	fmt.Println(strings.TrimSpace(" 123 "))

	// 去除字符串首尾指定的字符
	fmt.Println(strings.Trim("# 123 #", "#"))

	// 去除字符串右边指定的字符
	fmt.Println(strings.TrimRight("#123#", "#"))

	// 去除字符串左边指定的字符
	fmt.Println(strings.TrimLeft("#123#", "#"))

	// 返回指定字符串通过空格分隔的所有子串slice
	fmt.Println(strings.Fields("hello world hello golang"))

	// 返回指定字符串由指定字符分隔的所有子串slice
	fmt.Println(strings.Split("Hello,World,Hello,Golang", ","))

	// 用指定字符把thr中的所有元素连接起来组成字符串
	thr := []string{"hello", "world"}
	res := strings.Join(thr, "-")
	fmt.Println(res)
}
