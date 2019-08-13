package main

import "fmt"

func main() {
	one := []int{1, 2, 3, 4, 5}
	fmt.Println(one)
	one[0], one[4] = one[4], one[0]
	fmt.Println(one)
}
