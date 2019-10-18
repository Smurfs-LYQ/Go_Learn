package main

import "fmt"

func T1(a, b int) int {
	return a + b
}

func T2(a, b int) int {
	return a - b
}

func T3(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	Test_1 := T3(1, 2, T1)
	fmt.Println(Test_1)
	Test_2 := T3(1, 2, T2)
	fmt.Println(Test_2)
}
