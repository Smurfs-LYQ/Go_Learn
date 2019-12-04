package mylog

import (
	"fmt"
	"os"
	"path"
	"strings"
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
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0755)
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

// Debug 日志函数··
func (f *Filelog) Debug() {
	fmt.Println(f)
	fmt.Println("debug")
}

// Info 日志函数··
func (f *Filelog) Info() {

}

// Warning 日志函数··
func (f *Filelog) Warning() {

}

// Error 日志函数··
func (f *Filelog) Error() {

}

// Fatel 日志函数··
func (f *Filelog) Fatel() {

}
