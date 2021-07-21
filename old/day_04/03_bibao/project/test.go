package project

import (
	"fmt"
	"strings"
)

// 判断一个字符串是否以指定字符结尾
func end(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func Project_1() {
	t1 := end(".jpg")
	fmt.Println(t1("one"))
	fmt.Println(t1("two.jpg"))

	t2 := end(".png")
	fmt.Println(t2("one.png"))
	fmt.Println(t2("two"))
}
