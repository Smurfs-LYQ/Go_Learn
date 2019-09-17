package main

import "fmt"

// 使用type声明的是命名类型
type T1 struct {
	Name string
	Age int
}

func main() {
	// 使用struct字面量声明的是未命名类型
	a := struct {
		Name string
		Age int
	}{"Test_A", 18}
	fmt.Printf("类型: %T\n值: %v\n",a , a)

	b := T1{"Test_B", 20}
	fmt.Printf("类型: %T\n值:%v\n",b , b)
}
