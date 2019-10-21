package main

import "fmt"

func main() {
	var T1 string
	T1 = "Test_1"
	fmt.Println(T1)

	T2 := &T1 // 取变量T1的内存地址
	fmt.Printf("%T, %p\n", T2, T2)

	T3 := *T2 // 取变量T2的值(前提是变量T2的值必须是一个内存地址)
	fmt.Printf("%T, %s\n", T3, T3)

	Arr_1 := [3]int{1, 2, 3}
	Array_T1(Arr_1)
	fmt.Println(Arr_1)

	Array_T2(&Arr_1)
	fmt.Println(Arr_1)
}

func Array_T1(a [3]int) {
	a[0] = 100 // 只是修改的内部的a这个数组
}

func Array_T2(a *[3]int) {
	// (*a)[0] = 100 // 因为传入的是Arr_1这个数组的内存地址，所以修改的是Arr_1这个数组本身
	// 语法糖: 其实就是相对来说简化的写法
	a[0] = 100
}
