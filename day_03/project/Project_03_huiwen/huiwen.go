package Project_03_huiwen

import (
	"fmt"
)

// 输入一个字符串，判断其是否为回文。回文字符串是指从左到右读和从右到左读完全相同的字符串。例如：12121
func Project_huiwen() {
	var one string
	fmt.Scanf("%s\n", &one)

	str := []rune(one) // rune可以表示一个unicode字符(中文英文都可以), byte是表示一个字节可以存放的值(符号, 英文, 数字)
	len := len(str) - 1
	res := true
	for i, _ := range str {
		if str[i] == str[len] {
			len--
			continue
		} else {
			res = false
			break
		}
	}

	if res {
		fmt.Println(one, " 是回文")
	} else {
		fmt.Println(one, " 不是回文")
	}
}
