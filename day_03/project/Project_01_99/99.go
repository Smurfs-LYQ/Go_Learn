package Project_01_99

import "fmt"

// 输出99乘法表
func Project_99() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", j, i, i*j)
		}
		fmt.Println(" ")
	}
}
