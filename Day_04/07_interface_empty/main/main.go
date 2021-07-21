package main

import (
	"fmt"
)

// T1 声明一个可以接收任何类型参数的函数
func T1(a interface{}) {
	fmt.Printf("%T %v\n", a, a)
}

func main() {
	// 声明一个空接口
	var a interface{}
	// 因为任何类型都实现了空接口，所以任何类型都可以对空接口进行赋值
	a = 100
	fmt.Printf("%T %v\n", a, a)
	a = "Smurfs"
	fmt.Printf("%T %v\n", a, a)
	a = false
	fmt.Printf("%T %v\n", a, a)

	fmt.Println()

	// 调用可以接收任何类型参数的函数
	T1(123)
	T1("Smurfs")
	T1(false)

	fmt.Println()

	// 使用空接口实现可以保存任意值的字典
	var map_1 = make(map[int]interface{})
	map_1[0] = 123
	map_1[1] = struct{ name string }{name: "Smurfs"} // 传一个结构体也可以接收
	map_1[2] = false
	for k, v := range map_1 {
		fmt.Printf("%d %T %v\n", k, v, v)
	}
}
