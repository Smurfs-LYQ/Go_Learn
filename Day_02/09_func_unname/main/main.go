package main

import "fmt"

func main() {
	// 匿名函数定义之后加 () 直接调用
	func(x, y int) {
		fmt.Println(x + y)
	}(1, 2)

	// 将匿名函数赋值给变量
	T1 := func(x, y int) {
		fmt.Println(x - y)
	}
	T1(2, 1)
}
