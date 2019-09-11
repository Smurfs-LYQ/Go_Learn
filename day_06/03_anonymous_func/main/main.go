package main

import "fmt"

var sum = func (a, b int) int {
	return a+b
}

func doinput(f func(int, int) int, a, b int) int {
	return f(a, b)
}

func warp(op string) func(int, int) int {
	switch op {
	case "add" :
		return func (a, b int) int {
			return a+b
		}
	case "sub" :
		return func(a, b int) int {
			return a+b
		}
	default :
		return nil
	}
}

func main() {
	// defer 延迟调用, 这里没有其他条件, 所以直接调用
	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
	}() // 这里()的作用是调用匿名函数, ()里面可以传递参数

	sum(1, 2)

	// 传入参数：匿名参数, int, int
	res := doinput(func(x, y int) int {
		return x+y
	}, 1, 2)
	println(res)

	// 返回一个匿名函数, 然后赋值给opFunc变量
	opFunc := warp("add")
	// 给赋值给opFunc的匿名函数赋值
	re := opFunc(2, 3)

	fmt.Printf("%d\n", re)
}
