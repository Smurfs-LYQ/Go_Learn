package main

import (
	"fmt"
	"sync"
)

var x int64 // 定义一个全局变量
var wg sync.WaitGroup

// 定义一个互斥锁
var lock sync.Mutex

// 定义一个函数，对全局的变量x做累加操作
func add() {
	for i := 0; i < 1000; i++ {
		// 加锁
		lock.Lock()

		x++

		// 解锁
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
