package mylog

// 设置等级类型别名
type Level uint16

// 设置等级常量
const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatelLevel
)

type Mylogger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatel(format string, args ...interface{})
	Die()
}
