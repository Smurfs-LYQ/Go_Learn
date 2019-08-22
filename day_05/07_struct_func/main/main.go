package main

import "fmt"

/*
Golang中的方法是作用在特定类型的变量上，因此自定义类型，都可以有方法，而不仅仅是struct
	定义:
		func (recevier type) methodName(参数列表)(返回值列表){}
		recevier type : 用于标注这个函数属于哪个结构体
*/

type T1 struct {
	Name string
	Age  int
}

// 结构体中的函数
func (Test_1 *T1) init(name string, age int, sex string) {
	Test_1.Name = name
	Test_1.Age = age
	fmt.Println(Test_1, sex)
}

func (Test_2 T1) get() T1 {
	return Test_2
}

func main() {
	var one T1
	// (&one).init("test", 10, "未知")
	one.init("test", 10, "未知") // 效果同上，go会自动识别转换成(&one)

	fmt.Println(one.get())
}
