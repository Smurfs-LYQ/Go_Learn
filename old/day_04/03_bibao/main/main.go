package main

import (
	"Go_learn/old/day_04/03_bibao/project"
)

// 闭包: 一个函数和与其相关的引用环境组合而成的实体
/*
func Adder() func(int) int { // 定义一个函数Adder, 其返回值还是一个函数(这里定义的是这个函数的格式, 类似于type(类型)的作用)
	var x int
	return func(delta int) int { // 返回Adder函数的返回值(一个匿名函数)
		x += delta
		return x // Adder函数返回值 "(匿名)函数"的返回值
	}
}

func main() {
	var f = Adder()
	fmt.Println(f(1)) // 这里的f对应的就是Adder的返回值(函数), 因为这个函数需要一个整形实参, 所以我们需要传入一个整形 得到最后的返回值
	fmt.Println(f(20))
	fmt.Println(f(300))
}
*/

func main() {
	project.Project_1()
}
