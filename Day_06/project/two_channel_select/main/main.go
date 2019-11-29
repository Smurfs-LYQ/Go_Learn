package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var T1ch chan *T1
var T2ch chan *T2

type T1 struct {
	id  int
	num int
}

type T2 struct {
	*T1
	sum int
}

// Producter 生产者
func Producter(ch chan *T1) {
	var i int
	for {
		i++
		ch <- &T1{
			i,
			rand.Intn(1000),
		}
	}
}

// Consummer 消费者
func Consummer(T1ch chan *T1, T2ch chan *T2) {
	for res := range T1ch {
		T2ch <- &T2{
			res,
			calc(res.num),
		}
	}
}

func calc(num int) (sum int) {
	for num > 0 {
		sum += num % 10
		num = num / 10
	}

	return
}

func printResult(T2ch chan *T2) {
	for res := range T2ch {
		fmt.Printf("id: %d num: %d sum: %d\n", res.id, res.num, res.sum)
		time.Sleep(time.Second)
	}
}

func main() {
	T1ch = make(chan *T1, 100)
	T2ch = make(chan *T2, 100)

	go Producter(T1ch) // 生产随机数的goroutine
	for i := 0; i < 10; i++ {
		go Consummer(T1ch, T2ch) // 消费随机数的goroutine
	}

	// 获取用户输入的管道
	ch := make(chan bool, 1)

	go func() { // 获取键盘输入的goroutine
		/*
			// 这个代码的作用跟下面代码的作用是一样的
			var str string
			fmt.Scan(&str)
		*/
		os.Stdin.Read(make([]byte, 1))
		ch <- true
	}()

	go printResult(T2ch) // 打印结果的goroutine
	select {
	// 获取键盘输入通道中的值，如果获取到了程序就会退出
	case <-ch:
		return
	}
}
