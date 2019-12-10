package split

import "strings"

// Split 字符串切割函数
// 基准测试
/*
// 优化前
func Split(str, sep string) (res []string) {
	// 获取sep在str中首次出现的位置
	index := strings.Index(str, sep)

	for index >= 0 {
		res = append(res, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	res = append(res, str)
	return
}
*/

// 优化后
func Split(str, sep string) (res []string) {
	res = make([]string, 0, strings.Count(str, sep)+1) // strings.Count 用于判断sep在str中出现的次数
	index := strings.Index(str, sep)
	for index > -1 {
		res = append(res, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	res = append(res, str)
	return
}
