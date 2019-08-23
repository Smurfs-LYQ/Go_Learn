package main

import "fmt"

/*
Golang中的方法是作用在特定类型的变量上，因此自定义类型，都可以有方法，而不仅仅是struct
	定义:
		func (recevier type) methodName(参数列表)(返回值列表){}
		recevier type :
			recevier : 它相当于实例化结构体(var one T1)时候中的 one，
			type : 用于标注这个函数属于哪个结构体
*/

type T1 struct {
	Name string
	Age  int
}

// 结构体中的函数
func (Test_1 *T1) Init(name string, age int, sex string) { // 这里的Test_1代表的是下面的 (&one)
	Test_1.Name = name
	Test_1.Age = age
	fmt.Println(Test_1, sex)
}

func (Test_2 T1) Get() T1 {
	return Test_2
}

func main() {
	// 方法调用
	// 定义一个结构体
	var one T1

	// 通过结构体变量名调用方法
	// (&one).Init("test", 10, "未知")
	one.Init("test", 10, "未知") // 效果同上，go会自动识别转换成(&one)

	fmt.Println(one.Get())
}
