package main

import (
	"math/rand"
	"time"
)

// 生产者: 产生随机数
// 消费者: 计算每个随机数的每个位的数字的和

// 1个生产者 10个消费者

type T1 struct {
	id   int
	rand int
}

type T2 struct {
	T1
	num int
}

// Prodocer 生产者
func Prodocer() {

}

// Customer 消费者
func Customer() {
}

func main() {
	rand.Seed(time.Now().UnixNano())
}
