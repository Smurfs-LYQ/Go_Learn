package main

import "fmt"

/*
组合：
	如果一个struct嵌套了另一个有名结构体，那么这种模式就叫组合
*/

type T1 struct {
	Name string
	Age  int
}

func (obj *T1) Init(name string, age int) {
	(*obj).Name = name
	(*obj).Age = age
}

type T2 struct {
	Test_1 T1
	Sex    string
}

func main() {
	var one T2
	one.Test_1.Init("测试-1", 18) // 调用组合模式嵌套的结构体中的方法需要写上对应字段的变量名
	one.Sex = "未知"
	fmt.Println(one)
}
