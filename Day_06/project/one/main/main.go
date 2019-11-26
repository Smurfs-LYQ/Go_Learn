package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// 生产者: 产生随机数
// 消费者: 计算每个随机数的每个位的数字的和

// 1个生产者 10个消费者

func init() {
	rand.Seed(time.Now().Unix())
}

var wg sync.WaitGroup

// Prodocer 生产者
func Prodocer(ch chan int) {
	// defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- rand.Intn(1000)
	}
	close(ch)
}

// Customer 消费者
func Customer(ch chan int) {
	defer wg.Done()
	res := strconv.Itoa(<-ch)
	var num int
	for i := 0; i < len(res); i++ {
		a, _ := strconv.Atoi(string(res[i]))
		num += a
	}
	fmt.Printf("%s\t %d\n", res, num)
}

func main() {
	wg.Add(10)
	var ch1 = make(chan int)
	go Prodocer(ch1)

	var ch2 = make(chan int, 1)
	for res := range ch1 {
		ch2 <- res
		go Customer(ch2)
	}
	close(ch2)
	wg.Wait()
}
