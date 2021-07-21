package main

import (
	"fmt"
	"strings"
)

// 定义一个函数，接收一个string类型的变量，返回值是一个匿名函数
func T1(name string) func() {
	return func() {
		fmt.Println("hello", name)
	}
}

// 定义一个函数，接收一个string类型的变量，返回值是一个匿名函数(也接收一个string类型的变量，返回一个string)
func T2(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 定义一个函数，接收一个int类型的变量，返回值是两个匿名函数(两个匿名函数也都接收一个int类型的变量，各返回一个int)
func T3(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	del := func(i int) int {
		base -= i
		return base
	}
	return add, del
}

func main() {
	// #################T1 Begin#################

	/*
		Test_1就是一个闭包
		这个Test_1接收了T1返回的匿名函数
			因为匿名函数调用了他引用环境T1中的name变量(此闭包=该匿名函数+外层函数T1中的name变量)
	*/
	Test_1 := T1("Smurfs")
	Test_1() // 相当于执行了T1函数中的匿名函数

	// #################T1 Done#################

	// #################T2 Begin#################

	/*
		Test_2也是一个闭包
		因为Test_2接收函数T2返回的匿名函数，而这个匿名函数也需要一个string类型的参数，所以调用的时候也需要传一个string类型的参数进去
	*/
	Test_2 := T2(".txt")
	res := Test_2("Test_2")
	fmt.Println(res)

	// #################T2 Done#################

	// #################T3 Begin#################
	// Test_3_add, Test_3_del为闭包
	Test_3_add, Test_3_del := T3(100)
	res_add := Test_3_add(200) // base = 100 + 200
	res_del := Test_3_del(100) // base = 300 - 100
	fmt.Println(res_add, res_del)

	// #################T3 Done#################
}
