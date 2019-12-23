package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	// 随机数由一个Source生成。像Float64和Int这样的顶级函数使用默认共享的Source，他每次运行时产生的都是一系列确定的值，所以如果需要每次运行都想得到不同行为的值，需要使用Seed函数来初始化默认Source
	rand.Seed(time.Now().Unix())
}

func main() {
	// 生成10个100以内的随机数
	for i := 0; i < 10; i++ {
		// 生成100以内的随机数
		fmt.Println(rand.Intn(100))
	}

	fmt.Println("###########")

	// 生成10个随机复数
	for i := 0; i < 10; i++ {
		// 生成随机复数
		fmt.Println(rand.Float32())
	}
}
