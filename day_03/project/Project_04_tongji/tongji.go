package Project_04_tongji

import "fmt"

func detection(one int, two []int) bool {
	for _, v := range two {
		if one == v {
			return true
		}
	}
	return false
}

// 输入一行字符，分别统计出其中英文字母、空格、数字和其他字符的个数
func Project_tongji() {
	var str string = "123 hello"
	// fmt.Scanln(&str)

	letter_res := 0
	space_res := 0
	num_res := 0
	outher_res := 0

	// 英文字母的ASCII
	letter := []int{65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121}

	// 空格的ASCII
	space := []int{32}

	// 数字的ASCII
	num := []int{49, 50, 51, 52, 53, 54, 55, 56, 57}

	for i := 0; i < len(str); i++ {
		switch {
		case detection(int(str[i]), letter) == true:
			letter_res += 1
		case detection(int(str[i]), space) == true:
			space_res += 1
		case detection(int(str[i]), num) == true:
			num_res += 1
		default:
			outher_res += 1
		}
		/*
			if detection(int(str[i]), letter) {
				letter_res += 1
			} else if detection(int(str[i]), space) {
				space_res += 1
			} else if detection(int(str[i]), num) {
				num_res += 1
			} else {
				outher_res += 1
			}
		*/
	}

	fmt.Println("字母总数：", letter_res)
	fmt.Println("空格总数：", space_res)
	fmt.Println("数字总数：", num_res)
	fmt.Println("其他总数：", outher_res)
}

/*
func Project_tongji() {
	var str string
	fmt.Scanf("%s\n", &str)

	num := 0
	letter := 0
	space := 0
	outher := 0

	for i := 0; i < len(str); i++ {
		one := str[i : i+1]
		fmt.Printf("%T\n", one)
		if unicode.IsLetter(one) {
			// 判断是否为字母
			letter += 1
		} else if unicode.IsDigit(one) {
			// 判断是否为数字
			num += 1
		} else if unicode.IsSpace(one) {
			// 判断是否为空格
			space += 1
		} else {
			outher += 1
		}
	}
}
*/
