package assignmentDemo

import "fmt"

// One 赋值运算符
func One() {
	// = 赋值运算符
	var T1 int
	T1 = 123
	fmt.Println(T1)

	// += 相加后再赋值
	var T2 int
	T2 += 1 // T2++ 效果是一样的
	fmt.Println(T2)

	// -= 相减后再赋值
	var T3 int = 1
	T3 -= 1
	fmt.Println(T3)
}
