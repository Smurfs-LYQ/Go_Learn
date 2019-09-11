package main

import (
	"fmt"
)

func main() {
	one := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%T %d\n", one, one)
}
