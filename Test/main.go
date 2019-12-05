package main

import "fmt"

const (
	one int = 1 << iota
	two
	thr
	fou
)

func main() {
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(thr)
	fmt.Println(fou)
}
