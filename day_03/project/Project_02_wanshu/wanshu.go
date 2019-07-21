package Project_02_wanshu

import "fmt"

// 一个数如果恰好等于它的因子之和, 这个数就称为"完数"。例如6 = 1+2+3。编程找出1000以内所有的完数
func Project_wanshu() {
	for i := 2; i <= 1000; i++ {
		one := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				one += j
			}
		}
		if one == i {
			fmt.Println(i)
		}
	}
}
