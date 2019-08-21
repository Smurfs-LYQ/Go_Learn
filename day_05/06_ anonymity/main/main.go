package main

import (
	"fmt"
	"time"
)

/*
匿名字段
	1. 匿名字段只需要写需要的数据类型或结构体名就可以了
	2. 一种数据类型只能写一个，不然会报错，因为调用的时候是这种匿名字段的时候是直接调数据类型的
	3. 冲突字段
		* 结构体A 中调用了另一个结构体B做匿名字段，如果出现了字段冲突的情况，优先调用结构体A的字段
		* 结构体A 中调用了另外两个结构体B和结构体C做匿名字段，如果出现了结构体B和结构体C的字段冲突，那么就需要在调用字段的时候指定是哪个结构体(B/C)的
*/

type Cart struct {
	Name string
}

type Train struct {
	Cart  // 匿名字段 Cart结构体
	int   // 匿名字段 int数据类型
	start time.Time
}

func main() {
	var T1 Train
	// 访问匿名函数的方法有两种

	// No.1
	T1.Name = "Auti" // 直接调用匿名字段Cart结构体中的字段
	T1.start = time.Now()
	// T1.start = time.Now().Format("2006/02/01 15:04:05")

	// No.2
	T1.int = 4 // 只写了数据类型的字段可以直接调用其数据类型为其赋值

	fmt.Println(T1)
	fmt.Println(T1.start.Format("2006/01/02 15:04:05"))
}
