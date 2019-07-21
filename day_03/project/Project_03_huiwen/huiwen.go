package Project_03_huiwen

import (
	"fmt"
)

// 输入一个字符串，判断其是否为回文。回文字符串是指从左到右读和从右到左读完全相同的字符串。例如：12121
func Project_huiwen() {
	var str string
	fmt.Scanf("%s\n", &str)

	len := len(str) - 1
	res := true
	for i, _ := range str {
		if str[i] == str[len] {
			len -= 1
			continue
		} else {
			res = false
			break
		}
	}

	if res {
		fmt.Println(str, " 是回文")
	} else {
		fmt.Println(str, " 不是回文")
	}
}
