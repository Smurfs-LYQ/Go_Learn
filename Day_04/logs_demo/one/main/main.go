package main

import (
	"Go_Learn/Day_04/logs_demo/one/mylog"
	"fmt"
	"time"
)

func main() {
Begin:
	str := "这是一条测试日志"
	log := mylog.NewFileLoger("../logs/", "test.log")

	for {
		if !log.Debug("[%s] [%s]\n", "Debug", str) {
			fmt.Println("重新开始")
			goto Begin
		}
		// 睡眠1秒
		// time.Sleep(time.Second)
		// 定时器，每间隔1秒执行一次，相比Sleep效率更高
		time.Tick(time.Second)
	}
}
