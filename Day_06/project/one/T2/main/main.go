package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生产者: 产生随机数
// 消费者: 计算每个随机数的每个位的数字的和

// 1个生产者 10个消费者

var T1chan chan *T1
var T2chan chan *T2

type T1 struct {
	id  int
	num int
}

type T2 struct {
	*T1
	sum int
}

// Producer 生产者
func Producer(ch chan *T1) {
	// 设置只设置生成5次
	// for i := 0; i < 5; i++ {

	// 一直生成
	var id int
	for {
		id++
		// 生成随机数
		num := rand.Intn(1000)

		// 把结构体发送到通道中
		ch <- &T1{
			id:  id,
			num: num,
		}
	}
	// 关闭通道
	close(ch)
}

// Consumer 消费者
func Consumer(T1ch chan *T1, T2ch chan *T2) {
	/*
		// 消费者会一直取，直到T1ch管道中为空
		for T1 := range T1ch {
			T2ch <- &T2{
				T1,
				calc(T1.num),
			}
		}
	*/

	// /*
	// 一个消费者只取一个
	var T1 = <-T1ch

	T2ch <- &T2{
		T1,
		calc(T1.num),
	}
	// */
}

// calc 求和函数
func calc(num int) (sum int) {
	// 循环条件: num大于0
	for num > 0 {
		// 通过取余操作拿到数字最后一位
		sum += num % 10
		// num除10，使其丢掉最后一位数字，因为都是int类型，所以除出来的小数也会被舍弃
		num = num / 10
	}
	return
}

// 打印
func printResult(ch chan *T2) {
	for res := range ch {
		fmt.Printf("id:%d num:%v\tsum:%v\n", res.T1.id, res.T1.num, res.sum)
		if len(ch) == 0 {
			close(ch)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	T1chan = make(chan *T1, 100)
	T2chan = make(chan *T2, 100)

	go Producer(T1chan)
	for i := 0; i < 5; i++ {
		go Consumer(T1chan, T2chan)
	}
	printResult(T2chan)
}
