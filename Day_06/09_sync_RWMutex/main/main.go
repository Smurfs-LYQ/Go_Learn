package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// 读比写多的时候使用读写锁 能够提高性能

var x int64
var wg sync.WaitGroup

// var lock sync.Mutex     // 定义互斥锁
var rwlock sync.RWMutex // 读写互斥锁: 并发读加的是读锁，并发写加的是写锁

func write() {
	defer wg.Done()
	// 加锁
	// lock.Lock()

	// 加写锁
	rwlock.Lock()
	x++
	time.Sleep(time.Millisecond * 10)
	// 解写锁
	rwlock.Unlock()

	// 解锁
	// lock.Unlock()
}

func read() {
	defer wg.Done()
	// 加锁
	// lock.Lock()

	// 加读锁
	rwlock.RLock()
	fmt.Println(x)
	file, _ := os.OpenFile("./a.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	fmt.Fprintln(file, x)
	time.Sleep(time.Millisecond * 500)
	// 解读锁
	rwlock.RUnlock()

	// 解锁
	// lock.Unlock()
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
