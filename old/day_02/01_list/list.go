package main

import "fmt"

func list(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("%d+%d=%d\n", i, n-i, n) // %d是整形的占位符
	}
}

func main() {
	list(10)
}
