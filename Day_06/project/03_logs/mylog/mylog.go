package mylog

// Level 设置等级类型别名
type Level uint16

// 设置等级变量
const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatelLevel
)

// Mylogger 设置日志接口
type Mylogger interface {
	Debug()
	Info()
	Warning()
	Error()
	Fatel()
}
