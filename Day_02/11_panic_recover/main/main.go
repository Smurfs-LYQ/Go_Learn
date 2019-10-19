package main

import "fmt"

func T1() {
	fmt.Println("Test_1")
}

func T2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("T2 确实有错误")
		}
	}()
	// 在可能触发panic的代码之前加revover
	panic("panic in T2")
}

func T3() {
	fmt.Println("Test_3")
}

func main() {
	T1()
	T2()
	T3()
}
