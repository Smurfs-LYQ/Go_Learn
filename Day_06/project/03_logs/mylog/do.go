package mylog

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

// Filelog 日志结构体
type Filelog struct {
	level    Level
	fileName string
	filePath string
	file     *os.File
	maxSize  int64
}

// NewFilelog 构造函数
func NewFilelog(level, filename, filepath string, maxsize int64) *Filelog {
	return &Filelog{
		level:    levelInt(level),
		fileName: filename,
		filePath: filepath,
		file:     open(path.Join(filepath, filename)),
		maxSize:  maxsize,
	}
}

// open 打开文件
func open(filepath string) *os.File {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("文件打开失败，失败原因%v\n", err))
	}
	return file
}

// levelInt 返回等级的int类型
func levelInt(level string) Level {
	switch strings.Title(level) {
	case "Debug":
		return DebugLevel
	case "Info":
		return InfoLevel
	case "Warning":
		return WarningLevel
	case "Error":
		return ErrorLevel
	case "Fatel":
		return FatelLevel
	default:
		return DebugLevel
	}
}

// levelString 返回等级的string类型
func levelString(level Level) string {
	// fmt.Println(level)
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

// GetCallerInfo 获取正在运行函数的文件信息
func GetCallerInfo(skip int) (filename, funcname string, line int) {
	pc, filename, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}

	// 通过pc名拿到函数名
	funcname = path.Base(runtime.FuncForPC(pc).Name())
	filename = path.Base(filename)
	return
}

func (f *Filelog) log(level Level, format string, args ...interface{}) {
	// 判断等级
	if f.level > level {
		return
	}

	// 判断文件大小
	file, _ := f.file.Stat()
	if file.Size() >= f.maxSize {
		f.file.Close()

		name := fmt.Sprintf("%s.log", time.Now().Format("20060102150405"))
		os.Rename(path.Join(f.filePath, f.fileName), path.Join(f.filePath, name))

		f.file = open(path.Join(f.filePath, f.fileName))
	}

	// 写入文件 日志格式: [时间] [文件:行号] [函数名] [日志级别] [信息]
	filename, funcname, line := GetCallerInfo(3)
	str := fmt.Sprintf("[%s] [%s:%d] [%s] [%s] [%s]", time.Now().Format("20060102150405.000"), filename, line, funcname, levelString(level), fmt.Sprintf(format, args...))

	fmt.Println(str)

	_, err := fmt.Fprintln(f.file, str)
	if err != nil {
		fmt.Println(err)
	}
}

// Debug 日志函数··
func (f *Filelog) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

// Info 日志函数··
func (f *Filelog) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)
}

// Warning 日志函数··
func (f *Filelog) Warning(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)
}

// Error 日志函数··
func (f *Filelog) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)
}

// Fatel 日志函数··
func (f *Filelog) Fatel(format string, args ...interface{}) {
	f.log(FatelLevel, format, args...)
}
