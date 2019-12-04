package mylog

import "os"

// Filelog 日志结构体
type Filelog struct {
	level    Level
	fileName string
	filePath string
	file     *os.File
	maxSize  int64
}

// NewFilelog 构造函数
func NewFilelog() {

}
