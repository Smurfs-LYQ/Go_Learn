package two

import "fmt"

// Two 求所有数组元素的和
func Two() {
	var arr = [3]int{1, 2, 3}
	var res int

	for _, v := range arr {
		res += v
	}
	fmt.Println(res)
}
