package logicDemo

import "fmt"

// One 逻辑运算符
func One() {
	fmt.Println("----------逻辑运算符----------")

	a := true
	b := false
	fmt.Printf("%t && %t: %t\n", a, b, a && b)
	fmt.Printf("%t || %t: %t\n", a, b, a || b)
	fmt.Printf("!%t: %t\n", a, !a)
}
