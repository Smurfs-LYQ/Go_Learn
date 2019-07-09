package main

import "fmt"

func main() {
	var i, j, k int
	i = 153 % 10
	j = (153 / 10) % 10
	k = (153 / 100) % 10
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}
