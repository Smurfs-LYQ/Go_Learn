package mylog

import "os"

// Filelog 日志结构体
type Filelog struct {
	level    Level    // 日志级别
	fileName string   // 文件名
	filePath string   // 日志路径
	file     *os.File // 日志对象
}

// NewFilelog Filelog结构体的构造函数
func NewFilelog() {

}
