package main

import (
	"fmt"
	"time"
)

// 声明常量
const (
	Man    = 1
	Female = 2
)

func main() {
	// 获取当前秒数的时间戳
	secound := time.Now().Unix()

	if secound%Female == 0 {
		fmt.Println("Female")
	} else {
		fmt.Println("Man")
	}
}
