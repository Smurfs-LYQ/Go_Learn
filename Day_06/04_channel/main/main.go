package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func recv(ch chan bool) {
	res := <-ch // 阻塞
	fmt.Println(res)
	wg.Done()
}

// 无缓冲通道和有缓冲通道

func main() {
	wg.Add(2)

	// 无缓冲通道(不指定容量)
	ch1 := make(chan bool)
	go recv(ch1)
	ch1 <- true
	fmt.Println(len(ch1), cap(ch1))

	// 有缓冲通道(给指定容量)
	ch2 := make(chan bool, 3)
	go recv(ch2)
	ch2 <- false
	fmt.Println(len(ch2), cap(ch2))
	wg.Wait()
}
