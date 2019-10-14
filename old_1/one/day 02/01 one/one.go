// 变量使用的三种方式
package main

import "fmt"

func main() {
	// 1: 指定变量类型，如果不赋值，使用默认值
	var one int
	fmt.Println("one = ", one)

	// 2: 根据值自行判断变量类型(类型推导)
	var two = 123
	fmt.Println("two = ", two)

	// 3: 省略var，注意( :=左侧的变量不应该是已经声明过的，否则会导致编译错误)
	// 下面的方式等价于 var thr string = "Smurfs"
	thr := "Smurfs"
	fmt.Println("thr = ", thr)
}