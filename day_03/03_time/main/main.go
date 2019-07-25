package main

import (
	"fmt"
	"time"
)

// 检测程序运行耗时
func one() {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("123")
}

func main() {
	/*
		time包的一些常量
			const (
				Nanosecond Duration = 1				// 纳秒
				Microsecond = 1000 * Nanosecond		// 微妙
				Millisecond = 1000 * Microsecond	// 毫秒
				Second      = 1000 * Millisecond	// 秒
				Minute      = 60   * Second			// 分
				Hour        = 60   * Minute			// 小时
			)
	*/

	// 获取当前时间
	fmt.Println(time.Now())

	// 获取当前月
	fmt.Println(time.Now().Month())

	// 获取当前天
	fmt.Println(time.Now().Day())

	// 获取当前分钟
	fmt.Println(time.Now().Minute())

	// 获取当前年份
	fmt.Println(time.Now().Year())

	// 格式化
	// Golang诞生时间:2006/02/1 15:04:05
	now := time.Now()

	fmt.Println(now.Format("02/1/2006 15:04:05"))
	fmt.Println(now.Format("2006/02/1 15:04:05"))
	fmt.Println(now.Format("2006/02/1"))

	start := time.Now().UnixNano()
	one()
	end := time.Now().UnixNano()
	fmt.Println((end - start) / 1000)

	// 获取当前时间戳(秒)
	fmt.Println(time.Now().Unix())
	// 获取当前时间戳(纳秒)
	fmt.Println(time.Now().UnixNano())
}
