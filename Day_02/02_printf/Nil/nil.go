package Nil

import "fmt"

// One 指针
func One() {
	T1 := 10
	fmt.Printf("%p\n", &T1)
	fmt.Printf("%#p\n", &T1)
}
