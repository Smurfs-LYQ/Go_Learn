package main

import (
	"fmt"
	"sync"
	"time"
)

// 读比写多的时候使用读写锁 能够提高性能

var x int64
var wg sync.WaitGroup

// var lock sync.Mutex
var rwlock sync.RWMutex // 读写互斥锁: 并发读加的是读锁，并发写加的是写锁

func write() {
	defer wg.Done()

	// 加写锁
	rwlock.Lock()

	x++
	time.Sleep(time.Millisecond * 1)

	// 解写锁
	rwlock.Unlock()
}

func read() {
	defer wg.Done()

	// 加读锁
	rwlock.RLock()

	fmt.Println(x)
	time.Sleep(time.Millisecond * 5)

	// 解读锁
	rwlock.RUnlock()
}

func main() {
	start := time.Now()

	// 写10次
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	// 读1000次
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("共耗时: %v\n", end.Sub(start)) // 相当于end-start的操作
	// fmt.Printf("共耗时: %v\n", end.Sub(start).Seconds()) // 相当于end-start的操作, 并将其转化成秒
}
