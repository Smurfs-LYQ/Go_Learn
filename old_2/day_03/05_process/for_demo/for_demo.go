package for_demo

import (
	"fmt"
	"strings"
)

func For_demo_1() {
	// No.1
	for i := 1; i <= 5; i++ {
		fmt.Println(strings.Repeat("*", i))
	}

	// No.2 死循环
	/*
		for {
			fmt.Println(123)
		}

		for true {
			fmt.Println(123)
		}
	*/

	// No.3 条件判断
	num_1 := 1
	for num_1 <= 5 {
		fmt.Println(num_1)
		num_1++
	}

	/*
		// No.4 for range 语句
		str := "hello world,中国"
		for i, v := range str {
			// i是下标, v是值
			fmt.Printf("index[%d] val[%c] len[%d]\n", i, v, len([]byte(string(v))))
		}
	*/

	/*
		// No.5 break continue
		for i, v := range str {
			if i == 5 {
				continue // 跳出当前循环, 继续下一次循环
			} else if i > 10 {
				break // 结束循环
			}
			fmt.Printf("index[%d] val[%c] len[%d]\n", i, v, len([]byte(string(v))))
		}
	*/

}

// No.6 label和goto语句
func For_demo_2() {
	/*
	// No.6_1
	Label1: // 这就是一个label点, 可以使用coninue, break和goto语句调用
	for i := 1; i <= 3; i++ {
		fmt.Println("#")
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue Label1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
	*/

	// No.6_2 // 通过label实现for循环
	i := 0
	Label1:
		fmt.Println(i)
		i++
		if i == 5 {
			return
		}
		goto Label1
}
