package split

import "strings"

// Split 字符串切割函数
func Split(str, sep string) (res []string) {
	// 获取字符串第一次出现的位置
	index := strings.Index(str, sep)
	for index >= 0 {
		res = append(res, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	res = append(res, str)
	return
}
