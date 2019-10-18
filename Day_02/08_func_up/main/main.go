package main

import "fmt"

// 定义全局变量
var num_1 = 1

func T1() {
	// 定义一个局部变量
	num_1 := 100
	fmt.Println("全局变量", num_1)
}

func main() {
	T1()

	// 语句块内定义一个局部变量
	/*
		这里的i就是语句块中的局部变量，只在该语句块中生效
	*/
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
}
