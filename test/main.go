package main

import (
	"fmt"
)

func main() {
	var a = 100
	var b chan int = make(chan int, 1)
	fmt.Println(a)
	fmt.Println(b)
}
