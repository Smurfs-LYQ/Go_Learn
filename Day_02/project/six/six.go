package six

import (
	"fmt"
	"strings"
)

/*
Six 分金币
50枚金币，分配给下面几个人: Matthew, Sarah, Augusts, Heidi, Emilie, Peter, Giana, Adriano, Araron, Elizabeth
分配规则如下:
	1. 名字中包含e或者E: 1枚金币
	2. 名字中包含i或者I: 2枚金币
	3. 名字中包含o或者O: 3枚金币
	4. 名字中包含u或者U: 4枚金币
计算出每个用户分到多少金币，以及剩下多少金币
*/
func Six() {
	num := 50
	list := map[string]int{
		"Matthew":   0,
		"Sarah":     0,
		"Augusts":   0,
		"Heidi":     0,
		"Emilie":    0,
		"Peter":     0,
		"Giana":     0,
		"Adriano":   0,
		"Araron":    0,
		"Elizabeth": 0,
	}

	for k, v := range list {
		if strings.Contains(k, "e") || strings.Contains(k, "E") {
			v += 1
			num -= 1
		}
		if strings.Contains(k, "i") || strings.Contains(k, "I") {
			v += 2
			num -= 2
		}
		if strings.Contains(k, "o") || strings.Contains(k, "O") {
			v += 3
			num -= 3
		}
		if strings.Contains(k, "u") || strings.Contains(k, "U") {
			v += 4
			num -= 4
		}

		fmt.Printf("%s: %d\n", k, v)
	}

	fmt.Println(num)
}
