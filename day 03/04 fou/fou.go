// golang的字符类型
package main

import (
	"fmt"
)

func main() {
	var a byte = 'a'
	var b byte = '0'

	// 当我们直接输出byte值的时候，其实输出的是对应的ASCII码值
	fmt.Println(a)
	fmt.Println(b)

	// 如果我们希望输出对应的字符，需要使用格式化输出
	fmt.Printf("a = %c\nb = %c\n", a, b)

	// 整数类型输出对应的ASCII码
	var c int = 'a'
	fmt.Printf("c=%d\n", c)

	// var d char = 97
	// fmt.Printf("d=%c", d)

}