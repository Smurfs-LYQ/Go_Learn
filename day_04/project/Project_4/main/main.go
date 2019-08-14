package main

import "fmt"

func do(one []int, foot int) {
	fmt.Println("one = ", one)
	fmt.Println("foot = ", foot)

	var top, down *int

	for i := 0; i < len(one); i++ {
		if one[i] > foot {
			top = &one[i]
			break
		}
	}

	for i := len(one) - 1; i >= 0; i-- {

		if one[i] == *top {
			one = append(one, foot)
			down = &one[len(one)-1]
			break
		} else if one[i] < foot {
			down = &one[i]
			break
		}
	}

	fmt.Println(one, *top, *down)
	*top, *down = *down, *top
}

// 实现一个快速排序
func main() {
	int1 := []int{3, 5, 8, 1, 2, 9, 4, 7, 6}
	// fmt.Println(int1)
	for i := 0; i < 2; i++ {

		do(int1[:len(int1)-1], int1[len(int1)-1])
		fmt.Println(int1)
	}

}
