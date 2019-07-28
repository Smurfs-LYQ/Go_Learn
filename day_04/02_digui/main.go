package main

import "fmt"

func calc(n int) int {
	if n == 1 {
		return 1
	}

	// 相当于做了一个阶乘功能，每一次调用递归调用都会保留相乘之后的结果，最后返回的是相加的值
	res := calc(n-1) * n // 这里的*n操作乘的是calc(n-1)的返回值，也就是说*是递归执行完最后执行的
	/*
		calc(5-1) * 5	// 24 * 5 = 120
			clac(4-1) * 4	// 6 * 4 = 24
				clac(3-1) * 3	// 2 * 3 = 6
					clac(2-1) * 2	// 1 * 2 = 2
						if 1 == 1 {
							return 1 // 加上后面的 *2 其实返回的是2
						}
	*/
	fmt.Println(n, res)
	return res
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
