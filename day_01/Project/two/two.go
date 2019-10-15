package two

import "fmt"

// One 99乘法表
func One() {
	for i := 1; i < 10; i++ {
		for o := 1; o <= i; o++ {
			fmt.Printf("%d * %d = %d\t", o, i, o*i)
		}
		fmt.Println()
	}

	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")

	for i := 9; i > 0; i-- {
		for o := 1; o <= i; o++ {
			fmt.Printf("%d * %d = %d\t", o, i, o*i)
		}
		fmt.Println()
	}
}
