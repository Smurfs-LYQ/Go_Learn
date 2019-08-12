package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	1. Golang中的包
		* golang目前有150个标准的包, 覆盖了几乎所有的基础库
		* golang.org有所有包的文档, 没事都翻翻
*/

// 互斥锁
var lock sync.Mutex

func testMutex() {
	var a map[int]int
	a = make(map[int]int)

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	rand.Seed(time.Now().Unix())

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			// 上锁
			lock.Lock()
			b[1] = rand.Intn(100)
			// 解锁
			lock.Unlock()
		}(a)
	}

	time.Sleep(time.Second)

	// 上锁
	lock.Lock()
	fmt.Println(a)
	// 解锁
	lock.Unlock()

	time.Sleep(time.Second)
}

// 读写锁
var rwlock sync.RWMutex

func testRWMutex() {
	var a map[int]int
	a = make(map[int]int)

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	rand.Seed(time.Now().Unix())

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			// 加入写锁
			rwlock.Lock()
			b[1] = rand.Intn(1000)
			// 解除写锁
			rwlock.Unlock()
		}(a)
	}

	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			// 加入读锁
			rwlock.RLock()
			fmt.Println(b)
			// 解除读锁
			rwlock.RUnlock()
		}(a)
	}

	time.Sleep(time.Second * 5)
}

func main() {
	testMutex()
	// testRWMutex()
}
