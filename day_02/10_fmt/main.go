package main

import "fmt"

func main() {
	// 结构体
	type Test struct {
		Name string
	}
	// 变量
	var one = Test{Name: "Golang"}

	// 普通占位符
	fmt.Println("###########普通占位符###########")
	// %v	相应值的默认格式。
	fmt.Printf("hello %v \n", one)
	// %+v 	打印结构体时，会添加字段名
	fmt.Printf("hello %+v \n", one)
	// %#v 	相应值的Go"语法"表示
	fmt.Printf("hello %#v \n", one)
	// %T 	输出指定变量的相应类型
	fmt.Printf("one type = \"%T\"\n", one)
	// %% 	输出百分号
	fmt.Printf("100%%\n")

	// 布尔占位符
	fmt.Println("###########布尔占位符###########")
	// %t 	true\false
	fmt.Printf("%t\n", true)

	// 整数占位符
	fmt.Println("###########整数占位符###########")
	// %b 	二进制表示指定值
	fmt.Printf("15 = %b\n", 15)
	// %o 	八进制表示指定值
	fmt.Printf("15 = %o\n", 15)
	// %d   十进制表示指定值
	fmt.Printf("15 = %d\n", 15)
	// %x 	十六进制表示，字母形式为小写a-f
	fmt.Printf("13 = %x\n", 13)
	// %X 	十六进制表示，字母形式为大写A-F
	fmt.Printf("13 = %X\n", 13)
	// %c 	相应的Unicode码点所表示的字符
	fmt.Printf("%c\n", 0x4E2D)
	// %U 	Uincode格式: U+1234, 等同于 "U+%04X"
	fmt.Printf("0x32D2 = %U\n", 0x32D2)
	// %q 	单引号围绕的字符字面值，由Go语法安全的转义
	fmt.Printf("%q\n", 0x4E2D)

	// 浮点数和复数的组成部分(实部和虚部)
	fmt.Println("######浮点数和复数的组成部分(实部和虚部)######")
	// %e 	科学计数法 字符小写
	fmt.Printf("13.3 = %e\n", 13.3)
	// %E 	科学计数法 字符大写
	fmt.Printf("13.3 = %E\n", 13.3)
	// %f 	有小数点而无指数
	fmt.Printf("13.3 = %f\n", 13.3)
	// %g 	根据情况选择 %e或%f 以产生更紧凑的(无末尾的0)输出
	fmt.Printf("13.3 = %g\n", 13.3)
	// %G 	根据情况选择 %E或%f 以产生更紧凑的(无末尾的0)输出
	fmt.Printf("13.3 = %G\n", 13.3)

	// 字符串与字节切片
	fmt.Println("#########字符串与字节切片#########")
	// %s 	输出字符串表示(string类型或[]byte)
	fmt.Printf("hello %s\n", one)
	fmt.Printf("hello %s\n", []byte("Golang"))
	// %q 	双引号围绕的字符串，由Go语法安全的转义
	fmt.Printf("hello %q\n", "World")
	fmt.Printf("hello %q\n", one)
	// %x 	十六进制，小写字母，每个字节两个字符
	fmt.Printf("Go = %x\n", "Go")
	// %X 	十六进制，大写字母，每个字节两个字符
	fmt.Printf("Go = %X\n", "Go")

	// 指针
	fmt.Println("#############指针#############")
	// % p 	十六进制表示，前缀0x
	fmt.Printf("%p\n", &one)
}
