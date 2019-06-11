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
}