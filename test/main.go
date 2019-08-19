package main

import "fmt"

func do(val []int) {
	if len(val) < 1 {
		return
	}

	one, head, left, right := val[0], 0, 1, len(val)-1
	for head < right {
		if one < val[left] {
			val[left], val[right] = val[right], val[left]
			right--
		} else {
			val[head], val[left] = val[left], val[head]
			head++
			left++
		}
	}

	do(val[:head])
	do(val[head+1:])
}

func main() {
	int1 := []int{3, 5, 8, 1, 2, 9, 4, 7, 6}

	do(int1[:])
	fmt.Println(int1)
}
