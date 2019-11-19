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
