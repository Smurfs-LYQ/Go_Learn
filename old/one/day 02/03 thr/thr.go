// 变量使用的注意事项
package main

import "fmt"

func main() {
	// 该区域的数据值可以在同一类型范围内不断变化
	var T1 int = 10
	T1 = 20
	T1 = 30
	fmt.Println(T1)


	// 变量在同一个作用域内不能重名
	/*
	T1 := "Smurfs"
	fmt.Println(T1)
	*/
}