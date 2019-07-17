package switch_demo

import (
	"fmt"
	"math/rand"
	"time"
)

func Switch_demo_1() {
	var num int
	fmt.Scanf("%d", &num)

	/*
		// No.1
		switch num {
			case 1 :
				fmt.Println("num is 1")
				fallthrough // 功能: 执行完case 1之后会继续向下执行case 2中的内容
			case 2 :
				fmt.Println("num is 2")
			case 3 :
				fmt.Println("num is 3")
			case 4, 5, 6, 7, 8, 9 :
				fmt.Println("num in 4-9")
			default :
				fmt.Println("num greater than 9, so num is :", num)
		}
	*/

	/*
		// No.2
		switch {
			case 0 <= num && num < 10 :
				fmt.Println("In Case One")
			case 10 <= num && 15 >= num :
				fmt.Println("In Case Two")
			default :
				fmt.Println("Done")
		}
	*/

	/*
		// No.3
		switch num := 10; {
			case num == 0 :
				fmt.Println("num is 0")
			case num == 10 :
				fmt.Println("num is 10")
			default :
				fmt.Println("done")
		}
	*/
}

func Switch_demo_2() {
	// 生成0的10的随机整数
	rand.Seed(time.Now().Unix())
	num := rand.Intn(10)

	for {
		//获取用户输入的值
		var user_num int
		fmt.Scanf("%d\n", &user_num)

		test := false
		switch {
		case num == user_num:
			fmt.Println("You did it")
			test = true
		case num > user_num:
			fmt.Println("小了")
		case num < user_num:
			fmt.Println("大了")
		}
		if test {
			break
		}
	}
}
