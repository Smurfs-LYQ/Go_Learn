package main

import "fmt"

// 定义全局变量
/*
var T1 int = 1
var T2 = 2
*/

/*
var (
	T1 = 1
	T2 = 2
)
*/

func main() {
	// var T1,  T2 int = 1, 2

	// var T1, T2 = 1, "2"

	// T1, T2 := 1, "Smurfs"

	var (
		T1 = 1
		T2 = "Smurfs"
	)
	fmt.Println(T1, T2)
}