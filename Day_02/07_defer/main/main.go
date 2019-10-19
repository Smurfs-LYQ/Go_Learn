package main

import "fmt"

func main() {
	fmt.Println("Start...") // 代码执行顺序 1 程序执行顺序 1
	defer fmt.Println(1)    // 代码执行顺序 2 程序执行顺序 5
	// panic("stop") // 测试代码执行顺序
	defer fmt.Println(2)   // 代码执行顺序 3 程序执行顺序 4
	defer fmt.Println(3)   // 代码执行顺序 4 程序执行顺序 3
	fmt.Println("Done...") // 代码执行顺序 5 程序执行顺序 2
}
