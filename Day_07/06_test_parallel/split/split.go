package split

import (
	"strings"
)

// 定义一个切割字符串的包

// Split 用sep分割s
func Split(s, sep string) (res []string) {
	// 获取字符在字符串中首次出现的位置
	index := strings.Index(s, sep)
	for index > 0 {
		res = append(res, s[:index]) // 通过切片获取这个字符出现位置前面的字符并将他们放入切片res中
		s = s[index+1:]
		index = strings.Index(s, sep)
	}
	res = append(res, s) // 将切割之后剩下的字符串尾巴加入到切片res中
	return
}
