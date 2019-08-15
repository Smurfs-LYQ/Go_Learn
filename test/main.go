package main

import "fmt"

func do(values []int) {
	if len(values) <= 1 {
		return
	}

	val, head, left, right := values[0], 0, 1, len(values)-1
	for head < right {
		if val < values[left] {
			values[left], values[right] = values[right], values[left]
			right--
		} else {
			values[head], values[left] = values[left], values[head]
			head++
			left++
		}
	}

	fmt.Println(values)
	do(values[:head])
	do(values[head+1:])
}

func main() {
	int1 := []int{3, 5, 8, 1, 2, 9, 4, 7, 6}

	do(int1[:])
	fmt.Println(int1)
}
