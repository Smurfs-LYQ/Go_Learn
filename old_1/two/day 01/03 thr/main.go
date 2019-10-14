package main

import "fmt"

func main() {
	// 声明一个int变量
	x := 1
	// 声明x的指针
	g := &x
	fmt.Println(*g)
	fmt.Println(g)
	fmt.Printf("%t\n", g)

	// 复习指针
	a := 123
	b := &a // 创建变量a的指针并赋值给b
	fmt.Println(b) // 打印指针的值(变量a在内存中的地址)
	fmt.Println(*b) // 打印内存地址所指向的值
}