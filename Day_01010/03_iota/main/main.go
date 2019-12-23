package main

import "fmt"

// iota 枚举
const (
	aa = iota // 0
	bb
	cc
	dd
)

// 使用 _ 忽略某些值
const (
	one = iota // 0
	two
	_
	fou
)

// 声明中间插队
const (
	T1 = iota // 0
	T2 = 100  // 100
	T3 = iota // 2
	T4        // 3
)
const T5 = iota // 0

// 定义数量级
const (
	_ = iota // iota = 0
	/*
		KB = 1的二进制格式向左移10位
		移动前: 00000000 00000001
		移动后: 00000100 00000000
	*/
	KB = 1 << (10 * iota) // iota = 1
	MB = 1 << (10 * iota) // iota = 2
	GB = 1 << (10 * iota) // iota = 3
	TB = 1 << (10 * iota) // iota = 4
	PB = 1 << (10 * iota) // iota = 5
)

// 多个iota定义在一行
const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
)

const (
	test1 = "123"
	test2
)

// 常量
func main() {
	fmt.Println(aa, bb, cc, dd)

	fmt.Println(one, two, fou)

	fmt.Println(T1, T2, T3, T4, T5)

	fmt.Println(KB, MB, GB, TB, PB)
	fmt.Printf("%b \n", KB)

	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e, f)
}
