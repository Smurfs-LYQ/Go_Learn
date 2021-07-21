package main

import "fmt"

func main() {
	// 定义一个int类型的指针
	var T1 *int
	// 初始化int类型的指针(因为int是值类型所以用new)
	T1 = new(int)
	fmt.Printf("%T, %p, %d\n", T1, T1, *T1)
	// 给T1赋值
	*T1 = 1
	fmt.Printf("%T, %p, %d\n", T1, T1, *T1)
}
