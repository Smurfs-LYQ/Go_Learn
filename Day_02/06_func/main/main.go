package main

import "fmt"

// T1 定义一个球两个数总和 和 差的例子
func T1(a int, b int) (int, int) {
	return a + b, a - b
}

// T2 定义一个接收可变参数的函数
func T2(a ...string) {
	fmt.Printf("%T, %s\n", a, a)
	for _, v := range a {
		fmt.Printf("hello %s\n", v)
	}
}

// T3 定义一个接收(一个固定参数和一个可变参数)的函数
/*
	固定参数和可变参数同时出现时，可变参数要放在最后
*/
func T3(a string, b ...int) {
	var res int
	for _, v := range b {
		res += v
	}
	fmt.Printf("%s: %d\n", a, res)
}

func main() {
	// 函数的调用
	fmt.Println(T1(2, 1))

	T2("1", "2", "3")

	T3("和", 1, 2, 3)
}
