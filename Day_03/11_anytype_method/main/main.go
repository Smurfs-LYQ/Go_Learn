package main

import "fmt"

type MyInt int

func (m *MyInt) T1() {
	fmt.Printf("%s: %p, %d\n", "有内存地址啥的", m, *m)
}

func main() {
	var one MyInt = 21
	fmt.Printf("%p, %d\n", &one, one)
	(&one).T1()
}
