package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// 整形
	fmt.Println("---------整形---------")

	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10   “占位符 %d 表示十进制整形数字”

	// 二进制
	fmt.Printf("%b \n", a) // 1010 “占位符 %b 表示二进制整形数字”

	// 八进制 以 0 开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77 “占位符 %o 表示八进制整形数字”

	// 十六进程 以 0x 开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // xx “占位符 %x 表示十六进制整形数字(小写)”
	fmt.Printf("%X \n", c) // XX “占位符 %X 表示十六进制整形数字(大写)”

	// 变量的内存地址
	fmt.Printf("%p \n", &a) // 0xc000094008 “占位符 %p 表示十六进制的内存地址”

	// 浮点型
	fmt.Println("---------浮点型---------")

	fmt.Printf("%f \n", math.MaxFloat32) // “占位符 %f 表示浮点数”
	fmt.Printf("%f \n", math.MaxFloat64)

	// 复数
	fmt.Println("---------复数---------")

	var com1 complex64
	com1 = 1 + 2i
	var com2 complex128
	com2 = 2 + 3i
	fmt.Println(com1)
	fmt.Println(com2)

	// 字符串
	fmt.Println("---------字符串---------")

	var str1 string = "hello world"
	var str2 string = `
	hello
	你好
	`
	fmt.Println(str1, str2)
	// 常用的字符串函数
	// 获取字符串的长度
	fmt.Println("str1 字符串的长度: ", len(str1))
	// 字符串拼接
	var str3 string = "hello"
	var str4 string = "world"
	fmt.Println(str3 + str4)
	// 字符串格式化
	var str5 string = fmt.Sprintf("%s %s", str3, str4)
	fmt.Println(str5)
	// 分割字符串
	fmt.Println(strings.Split(str5, " "))
	// 判断是否包含指定字符
	fmt.Println("判断字符串str5是否包含\"hallo\"字符串", strings.Contains(str5, "hello"))
	// 前缀/后缀判断
	fmt.Println("判断字符串str5是否以\"h\"开头", strings.HasPrefix(str5, "h"))
	fmt.Println("判断字符串str5是否以\" \"结尾", strings.HasSuffix(str5, " "))
	// 子串出现的位置
	fmt.Println("\"e\"第一次出现的位置: ", strings.Index(str5, "e"))
	fmt.Println("\"l\"最后一次出现的位置: ", strings.Index(str5, "l"))
	// join操作
	sl1 := []string{"hello", "world"}
	fmt.Println("使用指定的字符将sl1中的元素连接起来: ", strings.Join(sl1, "-"))

	// 字符
	fmt.Println("---------字符---------")

	s1 := "Golang"
	c1 := 'G' // ASCII码下占一个字节(8位 8bit)
	fmt.Println(s1, c1)
	s2 := "中国"
	c2 := '中' // UTF-8编码下一个中文占3个字节
	fmt.Println(s2, c2)

	// 针对byte类型的 字符串遍历
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c \n", s1[i]) // 占位符 %c 表示一个字符
	}

	fmt.Println()

	// 针对rune类型的 字符串遍历
	for _, v := range s2 {
		fmt.Printf("%c \n", v)
	}

	// 修改字符串
	a1 := "big"
	// 强制类型转换
	byteS1 := []byte(a1)
	byteS1[0] = 'p'
	fmt.Printf("字符串修改前: %s\n字符串修改后: %s\n", a1, string(byteS1))

	a2 := "白色的"
	// 强制类型转换
	runeS2 := []rune(a2)
	runeS2[0] = '黑'
	fmt.Printf("字符串修改前: %s\n字符串修改后: %s\n", a2, string(runeS2))
}
