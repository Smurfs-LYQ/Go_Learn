package mylogger

// Level 代表日志级别
type Level uint16

// 定义具体的日志级别常量
const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatelLevel
)

// Mylog 接口
type Mylog interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatel(format string, args ...interface{})
	Die()
}
