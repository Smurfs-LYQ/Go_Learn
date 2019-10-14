package if_demo

import (
	"fmt"
	"strconv"
)

func If_demo_1() {
	/*
		// No.1
		if true {
			fmt.Println("No.1_YES")
		}

		// No.2
		if false {
			fmt.Println("No.2_YES")
		} else {
			fmt.Println("No.2_No")
		}

		// No.3
		var one int
		fmt.Scanf("%d", &one)
		if one == 1 {
			fmt.Println("No.3_one")
		} else if one == 2 {
			fmt.Println("No.3_two")
		} else if one == 3 {
			fmt.Println("No.3_thr")
		} else {
			fmt.Println("No.3_done")
		}
	*/
}

func If_demo_2() {
	var str string
	fmt.Scanf("%s", &str)

	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("无法转换为整形")
	} else {
		fmt.Println("转换之后的内容是: ", num)
	}
}
