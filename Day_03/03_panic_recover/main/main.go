package main

import "fmt"

func T1() {
	defer func() {
		err := recover() // 捕获panic的异常信号，并尝试将函数从当前的异常状态中恢复
		if err != nil {
			fmt.Println("recover捕获到了panic的异常信号: ", err)
		}
	}()

	var a []int
	a[0] = 100 // 这里会panic，因为a变量并没有初始化
}

func main() {
	T1()

	fmt.Println("看看recover之后会不会执行")
}
