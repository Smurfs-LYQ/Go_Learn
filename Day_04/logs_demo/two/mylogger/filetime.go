package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogTime 日志文件结构体_可以设定最大文件大小
type FileLogTime struct {
	level    Level    // 日志级别
	fileName string   // 日志文件名
	filePath string   // 日志文件地址
	file     *os.File // 成功日志文件对象
	errfile  *os.File // 错误日志文件对象
	time     int64    // 最大日志文件大小
	maxTime  int64
}

// NewFileLogTime FileLogTime结构体的构造函数
func NewFileLogTime(level, fileName, filePath string, maxtime int64) *FileLogTime {
	f := &FileLogTime{
		level:    getLevelInt(level),
		fileName: fileName,
		filePath: filePath,
		time:     time.Now().Unix(),
		maxTime:  maxtime,
	}
	f.initFile()
	return f
}

// initFile 打开文件方法
func (f *FileLogTime) initFile() {
	logFile := path.Join(f.filePath, f.fileName)

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("文件打开失败, 错误信息: %s", err))
	}
	f.file = file

	errFile, err := os.OpenFile(path.Join(f.filePath, fmt.Sprintf("err%s", f.fileName)), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("文件打开失败, 错误信息: %s", err))
	}
	f.errfile = errFile
}

// 判断文件是否需要被切割
func (f *FileLogTime) checkTime() bool {
	return time.Now().Unix()-f.time >= 60*f.maxTime
}

// 文件切割函数
func (f *FileLogTime) fileSplit(file *os.File) *os.File {
	fileName := file.Name()                                                           // 拿到文件完整路径
	backupName := fmt.Sprintf("%s.%s", fileName, time.Now().Format("20060102150405")) // 设置备份文件名
	file.Close()                                                                      // 文件关闭
	os.Rename(fileName, backupName)                                                   // 文件改名
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)     // 重新创建文件
	if err != nil {
		panic(fmt.Sprintf("文件打开失败, 错误信息: %s", err))
	}
	return file
}

// Die 文件关闭
func (f *FileLogTime) Die() {
	f.file.Close()
	f.errfile.Close()
}

// log 写入日志执行参数
func (f *FileLogTime) log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}
	str := fmt.Sprintf(format, args...)
	// 日志格式: [时间] [文件:行号] [函数名] [日志级别] [信息]
	filename, line, funcname := GetCallerInfo(3)
	str = fmt.Sprintf("[%s] [%s:%d] [%s] [%s] [%s]", time.Now().Format("2006-01-02 15:04:05.000"), filename, line, funcname, getLevelStr(level), str)
	if f.checkTime() {
		f.file = f.fileSplit(f.file)
		f.errfile = f.fileSplit(f.errfile)
		f.time = time.Now().Unix()
	}
	fmt.Fprintln(f.file, str)
	if level >= ErrorLevel {
		fmt.Fprintln(f.errfile, str)
	}
}

// Debug 方法
func (f *FileLogTime) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

// Info 方法
func (f *FileLogTime) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)

}

// Warning 方法
func (f *FileLogTime) Warning(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)

}

// Error 方法
func (f *FileLogTime) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)

}

// Fatel 方法
func (f *FileLogTime) Fatel(format string, args ...interface{}) {
	f.log(FatelLevel, format, args...)
}
