package main

import (
	"fmt"
)

func one() {
	a := 1
	b := 2
	var c int

	c = a
	a = b
	b = c
	fmt.Println(a)
	fmt.Println(b)
}

func two() {
	a := 1
	b := 2

	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	one()
	fmt.Println("###")
	two()
}
