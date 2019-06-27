package main

import (
	"Go_Learn/day_01/05_goroute_example/goroute"
	"fmt"
)

func main() {
	// 声明一个管道
	var pipe chan int
	// 创建管道
	pipe = make(chan int, 1)
	// 并发调用
	go goroute.Add(1, 2, pipe)

	// 从管道里拿数据并赋值给res变量
	res := <-pipe
	fmt.Println(res)
}
