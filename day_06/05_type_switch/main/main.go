package main

import "fmt"

func main() {
	s := "Hello, 世界"
	var a []byte
	a = []byte(s)

	var b string
	b = string(a)

	var c []rune
	c = []rune(s)

	fmt.Printf("类型: %T\t值: %v\n", a, a)
	fmt.Printf("类型: %T\t值: %v\n", b, b)
	fmt.Printf("类型: %T\t值: %v\n", c, c)
}
