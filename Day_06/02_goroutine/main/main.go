package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("b:", i)
	}
}

func main() {
	// 设置程序运行时占用几个CPU逻辑核心数
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	go a()
	go b()
	// time.Sleep(time.Second)
	wg.Wait()
}
