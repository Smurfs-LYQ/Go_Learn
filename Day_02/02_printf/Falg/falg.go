package Falg

import "fmt"

// One 其他falg
func One() {
	T1 := "Smurfs的格格巫"

	// 字符串格式化打印
	fmt.Printf("%s\n", T1)

	// 字符串指定长度格式化打印
	fmt.Printf("%15s\n", T1)

	// 字符串指定长度格式化打印
	fmt.Printf("%-15s\n", T1)

	// 字符串指定长度并只保留 . 后面位数的字符格式化打印
	fmt.Printf("%15.7s\n", T1)

	// 字符串指定长度并只保留 . 后面位数的字符格式化打印
	fmt.Printf("%-15.7s\n", T1)

	// 字符串指定长度格式化打印，如果长度不足用0在头部填充
	fmt.Printf("%015s\n", T1)
}
