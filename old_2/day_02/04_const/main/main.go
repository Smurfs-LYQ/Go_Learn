package main

import "fmt"

func main() {
	// 声明常量
	const (
		a = iota // 0 它可以从0开始依次递增1
		b        // 1
		c        // 2
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
