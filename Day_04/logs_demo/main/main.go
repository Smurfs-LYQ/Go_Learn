package main

import (
	"Go_Learn/Day_04/logs_demo/mylog"
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
		time.Sleep(time.Second * 1)
	}
}
