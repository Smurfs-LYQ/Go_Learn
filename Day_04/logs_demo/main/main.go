package main

import (
	"Go_Learn/Day_04/logs_demo/mylog"
	"fmt"
)

func main() {
	log := mylog.NewFileLoger("../logs/", "test.log")
	// log.Debug("[%v] [%v] %s\n", time.Now().Format("2006-01-02 15:04:05.000"), "Debug", "这是一条测试日志")
	log.Debug("[%s] %s\n", "Debug", "这是一条测试日志")

	// log.Info("这是一条info的日志")

	fmt.Println("写入完成")
}
