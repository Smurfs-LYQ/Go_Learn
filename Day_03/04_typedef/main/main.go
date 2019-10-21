package main

import "fmt"

// NewInt 自定义类型
/*
	基于Go语言的int类型，创建一个NewInt类型
*/
type NewInt int

// MyInt 类型别名: 只存在代码编写过程中，代码编译之后根本不存在MyInt类型
type MyInt = int

func main() {
	// 自定义类型
	var Test_1 NewInt
	fmt.Printf("%T, %d\n", Test_1, Test_1)

	// 类型别名
	var Test_2 MyInt
	fmt.Printf("%T, %d\n", Test_2, Test_2)
}
