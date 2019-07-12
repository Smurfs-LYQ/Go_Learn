package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 把一个整数转换成字符串
	fmt.Printf("%T\n", strconv.Itoa(123))

	// 把一个字符串转成整数
	one, err1 := strconv.Atoi("123")
	fmt.Printf("%T\n", one)
	fmt.Println(err1)
}
