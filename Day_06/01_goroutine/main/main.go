package main

import (
	"fmt"
	"sync"
)

// 实例化sync包中的WaitGroup结构体(它里面有一个计数器)
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // 计数器-1, wg.Add()中的参数-1

	fmt.Println("hello Golang", i)
}

func main() {
	defer fmt.Println("程序结束")

	/*
		1. 创建一个goroutine
		2. 在新的goroutine中执行hello函数
	*/
	wg.Add(10) // 创建一个等待标签, 参数为: 需要等待的次数
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	/*
		wg.Add(1) // 创建一个等待标签, 计数器+1, 参数为: 需要等待的次数
		go hello(0)
	*/

	fmt.Println("hello world")
	// time.Sleep(time.Second) // 因为创建goroutine的时候程序就已经执行完了，所以需要加上这个sleep函数让程序等一会goroutine

	// 等hello执行完(执行hello函数的那个goroutine执行完)
	wg.Wait() // 阻塞，一直等待所有的goroutine结束。当wg.Add()中的参数减成0的时候才会执行wg.Wait
}
