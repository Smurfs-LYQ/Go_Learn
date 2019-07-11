package Project_1

import "fmt"

/*
Project1
判断101-200之间有多少质数并输出质数(只能被1和本身整除，不能被其他数整除)
*/

/*
// 方案1
func Project1() {
	for i := 101; i <= 200; i++ {
		for j := 2; j < i; j++ {
			one := float64(i) / float64(j)
			two := float64(int(float64(i) / float64(j)))
			if one != two {
				if i != j+1 {
					continue
				} else {
					fmt.Println(i)
				}
			} else {
				break
			}
		}
	}
}
*/

// 方案2
func Project1() {
	for j := 101; j <= 200; j++ {
		res := true
		for i := 2; i < j; i++ {
			if j%i == 0 {
				res = false
				break
			}
		}
		if res {
			fmt.Println(j)
		}
	}
}
