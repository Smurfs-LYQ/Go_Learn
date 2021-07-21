package mylog

// 自定义一个日志库，实现日志记录的功能
/*
日志分级
DEBUG   测试
TRACE   链路追踪/行为分析
INFO    普通错误
WARNING 警告
ERROR 	严重错误
FATEL	致命错误
*/

const (
	DEBUG = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATEL
)
