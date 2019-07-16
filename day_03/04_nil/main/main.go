package main

import (
	"fmt"
)


func test(str1 *string) {
	*str1 = "hello golang"
}

func main() {
	one := "hello world"

	// 获取一个变量的内存地址
	two := &one
	fmt.Println(two)

	// 通过变量的内存地址获取此变量的值
	fmt.Println(*two)

	// 传入变量one的内存地址(two), 通过one的内存地址来修改one的变量值
	test(two)
	fmt.Println(one)
}
