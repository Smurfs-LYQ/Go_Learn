package thr

import "fmt"

// Thr 找出数组中未指定值的两个元素的下标，比如从数组[1,3,5,7,8]中找出和为8的两个元素的下标 分别为(0,3)和(1,2)
func Thr() {
	var arr = [...]int{1, 3, 5, 7, 8}

	for k1, v1 := range arr {

		for i := k1 + 1; i < len(arr); i++ {
			if v1+arr[i] == 8 {
				fmt.Printf("(%d, %d) => (%d + %d)\n", k1, i, v1, arr[i])
			}
		}
	}
}
