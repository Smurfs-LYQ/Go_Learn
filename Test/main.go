package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}
	b1 := a1
	a1[1] = 20
	fmt.Println(a1)
	fmt.Println(b1)

	a2 := []int{1, 2, 3}
	b2 := a2
	a2[1] = 20
	fmt.Println(a2)
	fmt.Println(b2)

	var a3 = [...]int{}
	a3 = new([...]int, 3, 3)
	a3[0] = 1
	fmt.Println(a3)
}
