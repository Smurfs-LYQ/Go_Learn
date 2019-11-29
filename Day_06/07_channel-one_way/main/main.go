package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 单向通道

// Write 写入(设置只允许写入的通道)
func Write(ch chan<- int) {
	// 向通道中写入数据
	ch <- rand.Intn(1000)
}

// Read 读取(设置只允许读取的通道)
func Read(ch <-chan int) {
	// 从通道中读取数据
	fmt.Println(<-ch)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int, 1)

	Write(ch)
	Read(ch)
}
