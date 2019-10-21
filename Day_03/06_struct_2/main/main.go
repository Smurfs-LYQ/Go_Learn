package main

import "fmt"

// 定义一个结构体
type Test_1 struct {
	Name string
	Age  int
}

func main() {
	// 基本实例化
	var T1 Test_1
	T1.Name = "Smurfs"
	T1.Age = 21
	fmt.Println(T1)

	// 实例化结构体，但其为结构体类型指针
	// var T2 = &Test_1
	var T2 = new(Test_1)
	// 一般写法
	(*T2).Name = "Smurfs的格格巫"
	// 语法糖(简单写法)
	T2.Age = 21
	fmt.Println(T2)

	// 结构体初始化_1
	var T3 = Test_1{
		Name: "Smurfs",
		Age:  21,
	}
	fmt.Println(T3)

	// 结构体初始化_2
	var T4 = Test_1{
		"Smurfs",
		21,
	}
	fmt.Println(T4)
}
