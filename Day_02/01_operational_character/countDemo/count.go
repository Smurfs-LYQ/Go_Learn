package countDemo

import "fmt"

// One 算数运算符
func One() {
	fmt.Println("----------算数运算符----------")
	a := 1
	b := 2

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", b, a, b/a)
	fmt.Printf("%d %% %d = %d\n", b, a, b%a)
	a++
	fmt.Printf("a++: %d\n", a)
	b--
	fmt.Printf("b--: %d\n", b)
}
