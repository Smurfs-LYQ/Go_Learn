package main

import (
	"Go_Learn/Day_04/logs_demo/two/mylogger"
	"time"
)

var logger mylogger.Mylog

func main() {
	// 向日志文件中写日志
	// logger := mylogger.NewFileLogger("debug", "test.log", "../logs")
	// logger := mylogger.NewFileLogger("info", "golang.log", "../logs")

	// 向终端打印日志
	// logger := mylogger.NewConsoleLogger("debug")
	// logger := mylogger.NewConsoleLogger("info")

	// logger = mylogger.NewFileLogSize("dubug", "golang.log", "../logs")
	// defer logger.Die()

	logger = mylogger.NewFileLogTime("dubug", "golang.log", "../logs", 3)
	defer logger.Die()

	for {
		time.Sleep(time.Minute)
		logger.Debug("%s", "Debug日志")
		logger.Info("%s", "Info日志")
		logger.Warning("%s", "Warning日志")
		logger.Error("%s", "Error日志")
		logger.Fatel("%s", "Fatel日志")
	}

}
