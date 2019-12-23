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
	one = "Hello World, Hello Golang"
	one = strings.Replace(one, "Hello", "你好", 1)
	fmt.Println(one)

	// 字符串计数
	one = "Hello World, Hello Golang"
	fmt.Println(strings.Count(one, "Hello"))

	// 字符串重复指定次数
	one = "Hello "
	fmt.Println(strings.Repeat(one, 3))

	// 转为小写
	one = "HELLO"
	fmt.Println(strings.ToLower(one))

	// 转为大写
	one = "hello"
	fmt.Println(strings.ToUpper(one))

	// 去掉字符串首尾的空白字符
	one = " Hello World "
	fmt.Println(strings.TrimSpace(one))

	// 去掉字符串首尾指定的字符
	one = "#-Hello World-#"
	fmt.Println(strings.Trim(one, "#"))

	// 去掉字符串左边指定的字符
	one = "#Hello World#"
	fmt.Println(strings.TrimLeft(one, "#"))

	// 去掉字符串右边指定的字符
	one = "#Hello World#"
	fmt.Println(strings.TrimRight(one, "#"))

	// 返回指定字符串通过空格分割的所有子串的slice(类似于数组)
	one = "Hello World Hello Golang"
	fmt.Println(strings.Fields(one))

	// 返回指定字符串由指定字符分割的所有子串的slice(类似于数组)
	one = "Hello World, Hello Golang"
	fmt.Println(strings.Split(one, ","))

	// 用指定字符把two中的所有元素连接起来, two可以是slice
	two := []string{"Hello", "World"}
	res := strings.Join(two, "-")
	fmt.Println(res)
}
