package relationDemo

import "fmt"

// One 关系运算符
func One() {
	fmt.Println("----------关系运算符----------")
	a := 1
	b := 2

	fmt.Printf("%d == %d:\t %t\n", a, b, a == b)
	fmt.Printf("%d != %d:\t %t\n", a, b, a != b)
	fmt.Printf("%d > %d:\t %t\n", a, b, a > b)
	fmt.Printf("%d < %d:\t %t\n", a, b, a < b)
	fmt.Printf("%d >= %d:\t %t\n", a, b, a >= b)
	fmt.Printf("%d <= %d:\t %t\n", a, b, a <= b)
}
