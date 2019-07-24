package Project_05_dashu

import (
	"fmt"
	"strings"
)

func res(num1, num2 string) string {
	one := len(num1)
	two := len(num2)

	if one < two {
		num1, num2 = num2, num1
	}
	result := []byte(num1)
	for i := one - 1; i >= two; i-- {
		a := int(num1[i]) - 48
		b := int(num2[i-two]) - 48
		he := a + b
		if he >= 10 {
			fmt.Println(he, " >= 10")
		} else {
			result[i] = byte(he + 48)
		}
	}
	return string(result)
}

// 计算两个大数相加的和，这两个大数会超过int64的表示范围
func Project_dashu() {
	var a string
	fmt.Scanf("%s\n", &a)
	var b string
	fmt.Scanf("%s\n", &b)

	// num1 := []byte(strings.TrimSpace(a))
	// num2 := []byte(strings.TrimSpace(b))
	num1 := strings.TrimSpace(a)
	num2 := strings.TrimSpace(b)

	// res(num1, num2)
	fmt.Println(res(num1, num2))
}
