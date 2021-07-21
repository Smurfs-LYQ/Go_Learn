package fou

import (
	"fmt"
	"strings"
)

// Fou "hou do you do"中每个单词出现的次数
func Fou() {
	str := "hou do you do"
	var con = make(map[string]int, 5)
	for _, v := range strings.Split(str, " ") {
		// 检查指定的键是否存在于指定的map中
		_, ok := con[v]
		if ok {
			con[v]++
		} else {
			con[v] = 1
		}
	}
	fmt.Println(con)
}
