package main

import (
	"fmt"
)

// T1 定义一个可以接受任何类型的函数
func T1(a interface{}) {
	// 通过 类型断言 判断参数a的类型是否为int，并且返回一个转换为int之后的值
	v, ok := a.(int)
	if ok {
		fmt.Printf("%T %d\n", v, v)
	} else {
		fmt.Printf("抱歉只支持int类型, 已为您转换为int类型，转换之后的值为: %v\n", v)
	}
}

// T2 定义一个可以接受任何类型的函数
func T2(a interface{}) {
	// 通过 switch 来配合类型断言
	switch v := a.(type) { // 这里的"type"代表的就是类型
	case int:
		fmt.Printf("参数a为int类型，值为: %d\n", v)
	case string:
		fmt.Printf("参数a为string类型，值为: %s\n", v)
	case bool:
		fmt.Printf("参数a为bool类型，值为: %b\n", v)
	default:
		fmt.Println("无法解析您传入参数的类型")
	}
}

func main() {
	T1(100)
	T1("Smurfs")

	fmt.Println()

	T2(100)
	T2("Smurfs")
}
