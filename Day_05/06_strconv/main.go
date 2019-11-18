package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		a的典故
			这是C语言遗留下的典故。C语言中没有string类型而是用字符数组(array)表示字符串。所以 Iota 对很多C系的程序员很好理解
	*/
	// Atoi() 将字符串类型的整数转换为int类型
	res1, err := strconv.Atoi("18")
	fmt.Printf("%d %T %v\n", res1, res1, err)

	// Itoa() 函数用于将int类型数据转换为对应的字符串表示
	res2 := strconv.Itoa(18)
	fmt.Printf("%s %T\n", res2, res2)

	/*
		Parse系列函数
			Parse类函数用于转换字符串为给定类型的值: ParseBool()、ParseFloat()、ParseInt()、ParseUint()
	*/

	// ParseBool 返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、TRUE、FALSE；否则返回错误
	fmt.Println(strconv.ParseBool("1"))

	// ParseInt 返回字符串表示的整数值，接受正负号
	/*
		base指定进制(2到36), 最高只有36进制, 因为数字加字母一共就36为, 如果base为0，则会从字符串前置判断，“0x”是16进制，“0”是8进制，否则是10进制；
		bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64分别代表int、int8、int16、int32、int64；
		返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax; 如果结果超出类型范围err.Error = ErrRange。
	*/
	res3, err := strconv.ParseInt("2", 10, 64)
	fmt.Printf("%d %T %v\n", res3, res3, err)

	// ParseUint 它类似于 ParseInt 但不接受正负号，用于无符号整型
	res4, err := strconv.ParseUint("2", 10, 64)
	fmt.Printf("%d %T %v\n", res4, res4, err)

	// Format系列函数
	/*
		Format系列函数实现了将给定类型数据格式化为string类型数据的功能。
	*/
	// FormatBool 根据 参数(bool类型) 返回true或false
	res5 := strconv.FormatBool(true)
	fmt.Printf("%v %T\n", res5, res5)

	// FormatInt 返回参数1的base进制的字符串表示。base必须在2到36之间，结果中会使用小心字母'a'到'z'表示大于10的数字
	res6 := strconv.FormatInt(36, 32)
	fmt.Printf("%v %T\n", res6, res6)

	// FormatUint 是FormatInt的无符号整数版本
	res7 := strconv.FormatUint(36, 2)
	fmt.Printf("%v %T\n", res7, res7)

	// FormatFloat 函数将浮点数表示为字符串并返回, 具体参数作用访问Golang手册
	res8 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Printf("%v %T\n", res8, res8)

	// 其他
	// IsPrint 返回一个字符是否是可打印的，和 unicode.IsPrint 一样，参数必须是: 字母(广义)、数字、标点、符号、ASCII空格
	res9 := strconv.IsPrint(',')
	fmt.Println(res9)

	// CanBackquote 返回字符串s是否可以不被修改的表示为一个单行的、没有空格和tab之外控制字符的反引号字符串。
	fmt.Println(strconv.CanBackquote("123"), strconv.CanBackquote("\n"))
}
