// 多变量声明
package main

import "fmt"

// 定义全局变量
/*
var T1 = "test-1"
var T2 = "test-2"
var name = "Smurfs"
*/
// 一次性声明全局变量(也可以用于下面的局部变量)
var (
	T1 = "t1"
	T2 = "t2"
	name = "Smurfs"
)

func main() {
	// 一次性声明多个变量(同类型)
	var one, two, thr int
	fmt.Println("one = ", one, "two = ", two, "thr = ", thr)

	// 一次性创建多个变量(不同类型)
	var n1, n2, n3 = 1, "二", 3
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	// 一次性创建多个变量(不同类型，类型推导)
	m1, m2, m3 := 1, "二", 3
	fmt.Println("m1=", m1, "m2=", m2, "m3=", m3)

	// 调取全局变量
	fmt.Println("T1 = ", T1, "T2 = ", T2, "name = ", name)
}