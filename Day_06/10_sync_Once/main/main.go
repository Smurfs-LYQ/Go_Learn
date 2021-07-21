package main

import (
	"fmt"
	"sync"
)

func T1(a int) {
	fmt.Println(a)
}

// 闭包
func closer(x int) func() {
	return func() {
		T1(x)
	}
}

var onlyOnce sync.Once

func main() {
	T1 := closer(10)
	onlyOnce.Do(T1)
}
