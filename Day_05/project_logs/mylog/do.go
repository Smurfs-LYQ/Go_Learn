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
	level    Level    // 日志级别
	fileName string   // 文件名
	filePath string   // 日志路径
	file     *os.File // 日志对象
	maxSize  int64    // 文件最大大小
}

// NewFilelog Filelog结构体的构造函数
func NewFilelog(level string, fileName, filePath string, maxSize int64) *Filelog {
	loger := &Filelog{
		level:    levelInt(level),
		fileName: fileName,
		filePath: filePath,
		maxSize:  maxSize,
	}

	loger.file = loger.OpenFile()

	return loger
}

// OpenFile 打开文件函数
func (f *Filelog) OpenFile() *os.File {
	// 定义日志文件完整路径
	file_path := path.Join(f.filePath, f.fileName)

	file, err := os.OpenFile(file_path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("文件打开失败")
	}

	return file
}

// levelInt 将字符串格式的level转换成int格式
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

// levelString 将Level格式的level转换成string格式
func levelString(level Level) string {
	switch level {
	case DebugLevel:
		return "Debug"
	case InfoLevel:
		return "Info"
	case WarningLevel:
		return "Warning"
	case ErrorLevel:
		return "Error"
	case FatelLevel:
		return "Fatel"
	default:
		return "Debug"
	}
}

// GetCallerInfo 获取正在运行函数的文件名 行号 和 函数名
func GetCallerInfo(skip int) (filename, funcname string, line int) {
	pc, filename, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}

	// 通过pc拿到函数名
	funcname = path.Base(runtime.FuncForPC(pc).Name())
	filename = path.Base(filename)
	return
}

// Die 程序结束关闭文件
func (f *Filelog) Die() {
	f.file.Close()
}

// log 日志写入文件操作
func (f *Filelog) log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}

	// 判断日志文件大小
	file_stat, _ := f.file.Stat()
	if file_stat.Size() >= f.maxSize {
		// 关闭文件
		f.file.Close()

		old_name := path.Join(f.filePath, f.fileName)
		new_name := path.Join(f.filePath, fmt.Sprintf("%v%s", time.Now().Format("20060102150405"), f.fileName))

		// 旧文件重命名
		os.Rename(old_name, new_name)

		// 打开新文件并将其赋值给f.file
		f.file = f.OpenFile()
	}

	// 日志格式: [时间] [文件:行号] [函数名] [日志级别] [信息]
	filename, funcname, line := GetCallerInfo(3)
	log := fmt.Sprintf(format, args)
	log = fmt.Sprintf("[%s] [%s:%d] [%s] [%s] [%s]", time.Now().Format("2006-01-02 15:04:05.000"), filename, line, funcname, levelString(level), log)
	_, err := fmt.Fprintln(f.file, log)
	if err != nil {
		panic(fmt.Sprintf("在%s: 日志写入失败", time.Now().Format("2006-01-02 15:04:05.000")))
	}
}

// Debug debug级别错误
func (f *Filelog) Debug(format string, args ...interface{}) {
	// fmt.Println(f)
	f.log(DebugLevel, format, args...)
}

// Info Info级别错误
func (f *Filelog) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)
}

// Warning Warning级别错误
func (f *Filelog) Warning(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)
}

// Error Error级别错误
func (f *Filelog) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)
}

// Fatel Fatel级别错误
func (f *Filelog) Fatel(format string, args ...interface{}) {
	f.log(FatelLevel, format, args...)
}
