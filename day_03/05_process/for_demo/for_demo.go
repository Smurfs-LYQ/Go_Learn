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

	// No.4 for range 语句
	str := "hello world,中国"
	for i, v := range str {
		fmt.Printf("index[%d] val[%c] len[%d]\n", i, v, len([]byte(v)))
	}
}
