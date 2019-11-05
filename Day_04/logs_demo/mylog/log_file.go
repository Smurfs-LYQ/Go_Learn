package mylog

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

// FileLog 日志文件结构体
type FileLog struct {
	level       int
	logFilePath string
	logFileName string
	logFile     *os.File
}

// NewFileLoger FileLog结构体的构造函数
func NewFileLoger(logFilePath, logFileName string) *FileLog {
	// 初始化日志文件
	file, err := os.OpenFile(fmt.Sprintf("%s%s", logFilePath, logFileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		panic(fmt.Sprintf("文件打开失败，错误原因: %v", err))
	}

	return &FileLog{
		// level:       level,
		logFilePath: logFilePath,
		logFileName: logFileName,
		logFile:     file,
	}
}

// Debug 测试日志
func (f *FileLog) Debug(format string, args ...interface{}) {
	/*
		日志的格式
		时间 日志的级别 那个文件 哪一行 哪一个函数 错误日志的信息
	*/
	// 获取当前时间
	time := time.Now().Format("2006-01-02 15:04:05.000")
	// 获取文件地址 哪一行 哪一个函数
	pc, file, line, ok := runtime.Caller(1) // 0代表当前层调用 1代表当前层的上一层调用 2代表当前层的上上层调用
	if !ok {
		panic("获取环境信息失败")
	}
	file = path.Base(file)
	// 根据pc获取函数名
	funcName := path.Base(runtime.FuncForPC(pc).Name())
	format = fmt.Sprintf("[%s] [%s:%s] [%d] %s", time, file, funcName, line, format)
	_, err := fmt.Fprintf(f.logFile, format, args...)
	if err != nil {
		panic(fmt.Sprintf("无法写入, 错误信息: %v\n", err))
	}
}

// Info info日志
func (f *FileLog) Info(str string) {
	var log = fmt.Sprintf("%s\n", str)
	f.logFile.WriteString(log)
}
