package main

import "fmt"

func main() {
	// 通过非递归的方式实现斐波那契数列, 打印前100个数
	// 前两个数都是1后面的所有数都是前两个数的和
	var a [10]int
	a[0] = 1
	a[1] = 1
	for i := 2; i < 10; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	for _, v := range a {
		fmt.Println(v)
	}
}
