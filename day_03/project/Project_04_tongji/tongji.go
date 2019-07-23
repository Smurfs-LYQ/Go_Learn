package Project_04_tongji

import (
	"bufio"
	"fmt"
	"os"
)

/*
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
		// if detection(int(str[i]), letter) {
		// 	letter_res += 1
		// } else if detection(int(str[i]), space) {
		// 	space_res += 1
		// } else if detection(int(str[i]), num) {
		// 	num_res += 1
		// } else {
		// 	outher_res += 1
		// }
	}

	fmt.Println("字母总数：", letter_res)
	fmt.Println("空格总数：", space_res)
	fmt.Println("数字总数：", num_res)
	fmt.Println("其他总数：", outher_res)
}
*/

func detection(result string) (letter_res, space_res, num_res, outher_res int) {
	str := []rune(result)
	for _, v := range str {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			letter_res++
		case v == 32:
			space_res++
		case v >= '0' && v <= '9':
			num_res++
		default:
			outher_res++
		}
	}
	return
}

func Project_tongji() {
	// 从终端读取一行的内容
	reader := bufio.NewReader(os.Stdin) // bufio是带缓冲区的IO, NewReader是初始化一个读的实例, os.Stdin代表的是从标准的终端输入读取
	result, _, err := reader.ReadLine() // 这个是在读的实例中选择读一行的方法

	if err != nil {
		fmt.Println("read from console err: ", err)
		return
	}

	lr, sr, nr, or := detection(string(result))
	fmt.Println("字母总数：", lr)
	fmt.Println("空格总数：", sr)
	fmt.Println("数字总数：", nr)
	fmt.Println("其他总数：", or)
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
