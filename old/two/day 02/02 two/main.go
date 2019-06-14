package main

import "fmt"

// const (
// 	// a = iota
// 	// b = iota
// 	// c = iota
// 	// d = iota
// 	a = 1 << iota
// 	b = 2 << iota
// 	c = 3 << iota
// 	d = 1 << iota
// )

func main() {
	// fmt.Printf("%g\n%g\n%g\n%g\n", a, b, c, d)
	a := 1 << 2
	// a << 2
	fmt.Println(a)
}