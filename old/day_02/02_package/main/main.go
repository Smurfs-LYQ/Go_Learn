package main

import (
	// one two可用
	// "Go_learn/old/day_02/02_package/add"
	//thr可用(设置包别名)
	a "Go_learn/old/day_02/02_package/add"
	"fmt"
)

// One
/*
func main() {
	fmt.Println(add.Name)
	fmt.Println(add.Age)
}
*/

// Two
/*
func main() {
	// 初始化变量
	add.Test()

	fmt.Println(add.Name)
	fmt.Println(add.Age)
}
*/

// Thr
func main() {
	// 初始化变量
	a.Test()

	fmt.Println(a.Name)
	fmt.Println(a.Age)
}
