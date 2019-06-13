package main

import (
	"fmt"
)

// 声明常量，iota可以看做自增的枚举变量(用来初始化常量)
const(
	one = iota
	two = iota
	thr = iota
)

func main() {
	fmt.Println(one, two, thr)
}