package main

import "fmt"

/*
继承：
	如果一个struct嵌套了另一个匿名结构体，那么这个结构可以直接访问匿名结构体的方法，从而实现了继承
多重继承
	如果一个struct继承了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的方法，从而实现了多重继承
*/

type T1 struct {
	Name string
	age  int
}

func (Test_1 T1) Get() (string, int) {
	return Test_1.Name, Test_1.age
}

type T2 struct {
	Name string
	age  int
	sex  string
}

func (Test_2 T2) T2_func() string {
	return "Test_2 多重继承"
}

type T3 struct {
	T1 // 和匿名字段类似，这里只需要写要继承的结构体名称就可以
	T2
}

func main() {
	// 定义一个结构体
	var one T3
	// 结构体中“继承的结构体”中字段的赋值
	one.T1.Name = "T1_Name"
	one.T1.age = 1
	one.T2.Name = "T2_Name"
	one.T2.age = 2
	one.sex = "未知"
	fmt.Println(one)
	// 调用父类结构体中的方法
	fmt.Println(one.Get())
	fmt.Println(one.T2_func())
	// 注意：如果继承的结构体中字段或方法发生了重复现象，只需要在初始化赋值的时候指定是哪个父结构体的字段或方法就可以了
}
