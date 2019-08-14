package main

import "fmt"

// 实现一个插入排序
/*
插入排序的原理:
	首先拿出数列中的第一个数, 因为他前面没有别的数字, 所以不需要对比
	然后再依次拿出数列中的所有数字, 每拿出一个数字, 都需要跟他前面所有的数字进行比较
	如果比前面的数字小就与前面的数字进行换位
*/
func main() {
	int1 := []int{10, 4, 2, 1, 3, 8, 5}

	for i := 0; i < len(int1); i++ { // 首先要依次拿到所有的数字one
		var key int = i // 因为Go语言自带的数字交换 交换完之后, 他们的key也变了, 所以这里创建一个key, 防止交换后key发生变化
		fmt.Println("############")
		for j := i - 1; j >= 0; j-- { // 依次获取数字one前面的所有数字two
			fmt.Printf("%d\t%d\t", int1[key], int1[j])
			if int1[key] < int1[j] { // 与数字one前面的所有数字two进行对比
				int1[key], int1[j] = int1[j], int1[key] // 如果数字two比数字one大, 那么进行位置的转换
			}
			key-- // 每对比一次, key - 1, 这样数字one和数字two更换了位置之后, 就不会因为键的原因导致获取不到正确的数字one
			fmt.Println(int1)
		}
	}

	// fmt.Println(int1)
}
