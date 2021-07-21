package Integer

import "fmt"

// One Integer (整形)
func One() {
	// Integer
	// %+d 	带符号的整形
	fmt.Printf("%+d\n", 21)

	// %q 	打印单引号
	fmt.Printf("%q\n", 21)

	// %b 	打印整形的二进制
	fmt.Printf("%b\n", 10)

	// %o 	不带0的八进制
	fmt.Printf("%o\n", 8)

	// %#o 	带0的八进制
	fmt.Printf("%#o\n", 8)

	// %x 	小写的十六进制
	fmt.Printf("%x\n", 123123)

	// %X 	大写的十六进制
	fmt.Printf("%X\n", 123123)

	// %#x 	带0x的十六进制
	fmt.Printf("%#x\n", 123123)

	// %U 	打印Unicode字符
	fmt.Printf("0x4E2D = %U\n", 0x4E2D)

	// %#U 	打印带字符的Unicode
	fmt.Printf("0x4E2D = %#U\n", 0x4E2D)

	// %c 	相应的Unicode码点所表示的字符
	fmt.Printf("%c\n", 0x4E2D)

	// Integer width
	// %5d 	整形长度为5，右对齐，左边留白
	fmt.Printf("%5d\n", 21)

	// %-5d 左对齐右边留白
	fmt.Printf("%-5d\n", 21)

	// %05d 数字前面补零
	fmt.Printf("%05d\n", 21)
}
