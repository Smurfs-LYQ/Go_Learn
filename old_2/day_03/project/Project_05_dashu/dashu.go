package Project_05_dashu

import (
	"fmt"
	"strings"
)

func res(num1, num2 string) (result string) {
	index1 := len(num1) - 1
	index2 := len(num2) - 1
	left := 0

	for index1 >= 0 && index2 >= 0 {
		c1 := int(num1[index1] - '0')
		c2 := int(num2[index2] - '0')

		sum := c1 + c2 + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}

		c3 := sum%10 + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
		index2--
	}

	for index1 >= 0 {
		c1 := int(num1[index1] - '0')

		sum := c1 + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}

		c3 := sum%10 + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
	}

	for index2 >= 0 {
		c2 := int(num2[index2] - '0')

		sum := c2 + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}

		c3 := sum%10 + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index2--
	}

	if left == 1 {
		result = fmt.Sprintf("1%s", result)
	}

	return
}

// 计算两个大数相加的和，这两个大数会超过int64的表示范围
func Project_dashu() {
	var a string
	fmt.Scanf("%s\n", &a)
	var b string
	fmt.Scanf("%s\n", &b)

	num1 := strings.TrimSpace(a)
	num2 := strings.TrimSpace(b)

	fmt.Println(res(num1, num2))
}
