package main

import "fmt"

// 实现一个选择排序
func main() {
	int1 := []int{10, 4, 2, 1, 3, 8, 5}
	var a, b int
	for i := 0; i < len(int1); i++ {
		a, b = int1[i], i

		for j := i + 1; j < len(int1); j++ {
			if a > int1[j] {
				a = int1[j]
				b = j
			}
		}

		int1[i], int1[b] = int1[b], int1[i]
	}
	fmt.Println(int1)
}
