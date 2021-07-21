package mylogger

import (
	"path"
	"runtime"
	"strings"
)

// 存放一些公用的工具函数

// GetCallerInfo 获取文件名，行号，函数名等信息
func GetCallerInfo(skip int) (fileName string, Line int, funcName string) {
	pc, fileName, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	// 根据pc拿到函数名
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	// 从fileName(文件全路径)中剥离处文件名
	fileName = path.Base(fileName)
	return fileName, line, funcName
}

// 根据字符串信息返回日志级别
func getLevelInt(levelstr string) Level {
	levelstr = strings.ToLower(levelstr) // 将字符串转换为纯小写
	switch levelstr {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warning":
		return WarningLevel
	case "error":
		return ErrorLevel
	case "fatel":
		return FatelLevel
	default:
		return DebugLevel
	}
}

// 根据日志级别返回对应的字符串信息
func getLevelStr(level Level) string {
	switch level {
	case 0:
		return "Debug"
	case 1:
		return "Info"
	case 2:
		return "Warning"
	case 3:
		return "Error"
	case 4:
		return "Fatel"
	default:
		return "Debug"
	}
}
