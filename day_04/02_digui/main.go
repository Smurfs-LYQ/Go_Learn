package main

import "fmt"

func calc(n int) int {
	if n == 1 {
		return 1
	}

	// 相当于做了一个阶乘功能，每一次调用递归调用都会保留相乘之后的结果，最后返回的是相加的值
	return calc(n-1) * n
}

func main() {
	/*
		fmt.Println(time.Now().Format("2006/01/02 15:04:05"))
		time.Sleep(time.Millisecond * 1000)
		// 递归 自己调用自己
		main()
	*/
	n := calc(5)
	fmt.Println(n)
}
