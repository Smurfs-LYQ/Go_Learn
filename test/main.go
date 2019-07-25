package main

import "fmt"

func test(one ...int) {
	for _, v := range one {
		fmt.Println(v)
	}
}

func main() {
	test(1, 2, 3, 4, 5)
}
