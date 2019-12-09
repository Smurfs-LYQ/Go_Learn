package split

import "strings"

// Split 自己写的切割
func Split(str, set string) (res []string) {
	// 获取set第一次在str中出现的位置
	index := strings.Index(str, set)
	for index >= 0 {
		res = append(res, str[:index])
	}
	return
}
