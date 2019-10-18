package main

import "fmt"

func main() {
	fmt.Println("Start...") // 执行顺序 1
	defer fmt.Println(1)    // 执行顺序 5
	defer fmt.Println(2)    // 执行顺序 4
	defer fmt.Println(3)    // 执行顺序 3
	fmt.Println("Done...")  // 执行顺序 2
}
