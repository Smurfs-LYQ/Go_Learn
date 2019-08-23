package main

import "fmt"

/*
接口演示
	在进行打印操作的时候会自动调用String方法
*/

type T1 struct {
	Name string
	Age  int
}

func (obj *T1) Init(name string, age int) {
	(*obj).Name = name
	obj.Age = age
}

func (obj T1) String() string {
	return fmt.Sprintf("name=[%s]\nage =[%d]", obj.Name, obj.Age)
}

func main() {
	var one T1

	// (&one).Init("Test_1", 18)
	one.Init("Test_1", 18) // 效果同上，Go会自动调用
	fmt.Println(one)

	var Two T1
	fmt.Println(Two)
}
