package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogSize 日志文件结构体_可以设定最大文件大小
type FileLogSize struct {
	level    Level    // 日志级别
	fileName string   // 日志文件名
	filePath string   // 日志文件地址
	file     *os.File // 成功日志文件对象
	errfile  *os.File // 错误日志文件对象
	maxSize  int64    // 最大日志文件大小
}

// NewFileLogSize FileLogSize结构体的构造函数
func NewFileLogSize(level, fileName, filePath string) *FileLogSize {
	f := &FileLogSize{
		level:    getLevelInt(level),
		fileName: fileName,
		filePath: filePath,
		maxSize:  10 * 1024 * 1024,
	}
	f.initFile()
	return f
}

// initFile 打开文件方法
func (f *FileLogSize) initFile() {
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
func (f *FileLogSize) checkSize(file *os.File) bool {
	// 拿到一个文件的详细信息
	fileObj, _ := file.Stat()
	// 获取文件的大小
	return f.maxSize <= fileObj.Size()
}

// 文件切割函数
func (f *FileLogSize) fileSplit(file *os.File) *os.File {
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
func (f *FileLogSize) Die() {
	f.file.Close()
	f.errfile.Close()
}

// log 写入日志执行参数
func (f *FileLogSize) log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}
	str := fmt.Sprintf(format, args...)
	// 日志格式: [时间] [文件:行号] [函数名] [日志级别] [信息]
	filename, line, funcname := GetCallerInfo(3)
	str = fmt.Sprintf("[%s] [%s:%d] [%s] [%s] [%s]", time.Now().Format("2006-01-02 15:04:05.000"), filename, line, funcname, getLevelStr(level), str)
	if f.checkSize(f.file) {
		f.file = f.fileSplit(f.file)
	}
	fmt.Fprintln(f.file, str)
	if level >= ErrorLevel {
		if f.checkSize(f.errfile) {
			f.errfile = f.fileSplit(f.errfile)
		}
		fmt.Fprintln(f.errfile, str)
	}
}

// Debug 方法
func (f *FileLogSize) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

// Info 方法
func (f *FileLogSize) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)

}

// Warning 方法
func (f *FileLogSize) Warning(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)

}

// Error 方法
func (f *FileLogSize) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)

}

// Fatel 方法
func (f *FileLogSize) Fatel(format string, args ...interface{}) {
	f.log(FatelLevel, format, args...)
}
