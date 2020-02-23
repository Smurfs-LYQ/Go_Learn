package one

import (
	// 导入这个包但是并不会用到它，只是用它的初始化来打印字符串
	_ "Go_learn/old/day_02/03_init/test"
)

// Name 全局变量
var Name string

// Age 全局变量
var Age int

// 这个函数会被自动调用，但是并不能当做入口函数使用
func init() {
	Name = "GoLang"
	Age = 10
}
