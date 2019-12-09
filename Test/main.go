package main

import (
	"fmt"
)

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
	fmt.Printf("%c[%d;%d;%dm%s%c[0m\n", 0x1B, 41, 33, 0, "test", 0x1B)
	fmt.Printf("%c[%d;%d;%dm%s%c[0m\n", 0x1B, 41, 33, 1, "test", 0x1B)
	fmt.Printf("%c[%d;%d;%dm%s%c[0m\n", 0x1B, 41, 33, 4, "test", 0x1B)
	fmt.Printf("%c[%d;%d;%dm%s%c[0m\n", 0x1B, 41, 33, 5, "test", 0x1B)
	fmt.Printf("%c[%d;%d;%dm%s%c[0m\n", 0x1B, 41, 33, 7, "test", 0x1B)
	fmt.Printf("%c[%d;%d;%dm%s%c[0m\n", 0x1B, 41, 33, 8, "test", 0x1B)
}
