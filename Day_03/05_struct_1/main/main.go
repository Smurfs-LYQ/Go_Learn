package main

import "fmt"

// 定义一个结构体
type T1 struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	// 结构体初始化
	var Test_1 = T1{
		Name: "Smurfs",
		Age:  21,
		Sex:  "男",
	}
	fmt.Printf("%T, %v\n", Test_1, Test_1)
	// 调用结构体中的单个数据
	fmt.Println(Test_1.Name)
	fmt.Println(Test_1.Age)
	fmt.Println(Test_1.Sex)
}
