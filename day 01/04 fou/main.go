package main

import "fmt"

func main() {
	var x, y int
	fmt.Println(&x, "\n")
	fmt.Println(&y, "\n")
	fmt.Println("------------")
	fmt.Println(&x == &x, &x == &y, &x == nil)
}