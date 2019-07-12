package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 把一个整数转换成字符串
	fmt.Printf("%T\n", strconv.Itoa(123))
	// 也可以用字符串格式化的方式转换数据类型
	one := 123
	res := fmt.Sprintf("%d\n", one)
	fmt.Printf("%T\n", res)

	// 把一个字符串转成整数
	two, err1 := strconv.Atoi("123")
	fmt.Printf("%T\n", two)
	fmt.Println(err1)
}
