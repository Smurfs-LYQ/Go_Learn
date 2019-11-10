package main

import "Go_Learn/Day_04/logs_demo/two/mylogger"

var logger mylogger.Mylog

func main() {
	// 向日志文件中写日志
	// logger := mylogger.NewFileLogger("debug", "test.log", "../logs")
	// logger := mylogger.NewFileLogger("info", "golang.log", "../logs")

	// 向终端打印日志
	// logger := mylogger.NewConsoleLogger("debug")
	// logger := mylogger.NewConsoleLogger("info")

	logger = mylogger.NewFileLogSize("dubug", "golang.log", "../logs")
	defer logger.Die()

	for {
		logger.Debug("%s", "测试日志")
		logger.Info("%s", "测试日志")
		logger.Warning("%s", "测试日志")
		logger.Error("%s", "测试日志")
		logger.Fatel("%s", "测试日志")
	}

}
