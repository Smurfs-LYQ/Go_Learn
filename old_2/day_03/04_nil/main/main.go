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

	// 声明变量thr为字符串指针
	var thr *string
	// 给thr赋值，值为one的内存地址
	thr = &one
	// 修改thr的值，因为它指向的时one的内存，所以one也会发生改变
	*thr = "Hello LYQ"
	fmt.Println(one)
}
