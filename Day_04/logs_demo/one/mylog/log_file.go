package mylog

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

// FileLog 日志文件结构体
type FileLog struct {
	level       int
	logFilePath string
	logFileName string
	logFile     *os.File
	time        int64
}

// NewFileLoger FileLog结构体的构造函数
func NewFileLoger(logFilePath, logFileName string) *FileLog {
	// 初始化日志文件
	logFileName_1 := strings.Split(logFileName, ".")
	date := time.Now()
	if len(logFileName_1) > 1 {
		logFileName = fmt.Sprintf("%s.%v.%s", logFileName_1[0], date.Format("2006-01-02_15:04:05"), logFileName_1[1])
	} else {
		logFileName = fmt.Sprintf("%s.%v.log", logFileName_1[0], date.Format("2006-01-02_15:04:05"))
	}

	file, err := os.OpenFile(fmt.Sprintf("%s%s", logFilePath, logFileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		panic(fmt.Sprintf("文件打开失败，错误原因: %v", err))
	}

	return &FileLog{
		// level:       level,
		logFilePath: logFilePath,
		logFileName: logFileName,
		logFile:     file,
		time:        date.Unix(),
	}
}

// Debug 测试日志
func (f *FileLog) Debug(format string, args ...interface{}) bool {
	/*
		日志的格式
		时间 日志的级别 那个文件 哪一行 哪一个函数 错误日志的信息
	*/
	// 获取当前时间
	date := time.Now()
	/*
		// 判断时间切割 以分钟来进行切割
		if (date.Unix() - f.time) > 60 {
			f.logFile.Close()
			return false
		}
	*/
	// 通过文件大小切割
	if size, err := os.Stat(fmt.Sprintf("%s%s", f.logFilePath, f.logFileName)); err == nil {
		if size.Size()/1024/1024 >= 2 {
			f.logFile.Close()
			return false
		}
	} else {
		panic("获取文件信息失败")
		return false
	}
	time := date.Format("2006-01-02 15:04:05.000")
	// 获取文件地址 哪一行 哪一个函数
	pc, file, line, ok := runtime.Caller(1) // 0代表当前层调用 1代表当前层的上一层调用 2代表当前层的上上层调用
	if !ok {
		panic("获取环境信息失败")
	}
	file = path.Base(file)
	// 根据pc获取函数名
	funcName := path.Base(runtime.FuncForPC(pc).Name())
	format = fmt.Sprintf("[%s] [%s:%s] [%d] %s", time, file, funcName, line, format)
	fmt.Printf(format, args...)
	_, err := fmt.Fprintf(f.logFile, format, args...)
	if err != nil {
		panic(fmt.Sprintf("无法写入, 错误信息: %v\n", err))
	}
	return true
}

// Info info日志
func (f *FileLog) Info(str string) {
	var log = fmt.Sprintf("%s\n", str)
	f.logFile.WriteString(log)
}
