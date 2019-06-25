package main

import "fmt"

func main() {

	// 变量声明
	//标准格式化输出(手动初始化)
	var a string = "hello, world"
	//编译器推到类型的格式
	var b = 1.123
	// 标准格式化输出(自动初始化)
	var c int
	c = 100
	// 简短格式
	d := true

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println("#################")

	// 多变量声明
	// 标准格式化，也可用于声明常量
	var (
		e int
		// f string = "多变量声明"
		f = "多变量声明"
	)
	fmt.Println(e)
	fmt.Println(f)

	g, h, i := 123, "一二三", b+float64(c)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	fmt.Println("#################")

	j, k := 1, 2
	fmt.Println(j, k)

	j = j ^ k
	k = k ^ j
	j = j ^ k
	fmt.Println(j, k)

	j, k = k, j
	fmt.Println(j, k)
	j, k = k, j
	fmt.Println(j, k)
}
1