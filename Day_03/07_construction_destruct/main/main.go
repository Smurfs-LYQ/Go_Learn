package main

import "fmt"

// 定义一个结构体
type T1 struct {
	Name string
	Age  int
}

// 定义一个初始化结构体的方法
/*
创建T1类型的变量时，直接调用Test_1函数
*/
func Test_1(name string, age int) *T1 {
	return &T1{
		name,
		age,
	}
}

func main() {
	// 调用自定义析构函数
	one := Test_1("Smurfs", 21)
	fmt.Println(one)
}
