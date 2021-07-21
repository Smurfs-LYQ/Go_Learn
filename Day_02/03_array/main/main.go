package main

import "fmt"

func main() {
	// 声明一个数组
	var arr1 [5]int
	// 初始化
	arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	// 声明一个数组并初始化
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	// 声明一个长度自动检测数组并初始化
	var arr3 = [...]string{"one", "two", "thr"}
	fmt.Println(arr3)

	// 声明一个长度自动检测并且带下标的数组并初始化
	var arr4 = [...]string{0: "one", 1: "two", 2: "thr"}
	fmt.Println(arr4[0], arr4[1], arr4[2], arr4)

	// 使用for循环遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	// 使用for range循环遍历数组
	for k, v := range arr4 {
		fmt.Println(k, v)
	}

	fmt.Println("-----------------------多维数组-----------------------")

	// 多维数组
	arr_1 := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(arr_1)

	arr_2 := [...][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(arr_2)

	// 多维数组的遍历
	for k, v1 := range arr_1 {
		fmt.Printf("第%d行:\t", k)
		for _, v2 := range v1 {
			fmt.Printf("%d \t", v2)
		}
		fmt.Println()
	}
}
