package main

import (
	"fmt"
)

func main() {
	// var a int = 1
	a := 1
	// var b chan int = make(chan int, 1)
	b := make(chan int, 1)
	fmt.Println(a)
	fmt.Println(b)
}
