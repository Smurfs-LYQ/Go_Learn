package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type T1 struct {
	SL   string
	SPBM string
}

func main() {
	str := "{\"code\":\"0000\",\"data\":\"{\"KXZMXS\":[{\"SL\":\"0\",\"SPBM\":\"1070101010300000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101020200000000\"},{\"SL\":\"7549.90999\",\"SPBM\":\"1070101030100000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101010500000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101020100000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101030300000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101040200000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101050100000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101010200000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101060100000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101030400000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101040300000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101050200000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101050300000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101070100000000\"},{\"SL\":\"17028.76998\",\"SPBM\":\"1070101010100000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101070200000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101070300000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101030200000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101040100000000\"},{\"SL\":\"0\",\"SPBM\":\"1070101010400000000\"}]}\",\"message\":\"成功\"}"

	str = strings.Trim(str, "{}")
	res1 := strings.Split(str, "\":\"")

	// 截取第一段信息 状态码
	res1_1 := strings.Split(res1[1], "\",\"")[0]
	fmt.Println("code:", res1_1)
	// 截取第二段信息 报文信息
	var res1_2 = strings.TrimPrefix(strings.TrimSuffix(strings.TrimRight(strings.Join(res1[2:len(res1)-1], "\":\""), "\",\"message"), "]}"), "{\"KXZMXS\":[")
	fmt.Println("KXZMXS:", res1_2)
	// 截取第三段信息 查询状态
	res1_3 := strings.TrimRight(res1[len(res1)-1], "\"")
	fmt.Println("message:", res1_3)

	res1_2_1 := strings.Split(strings.Trim(res1_2, "{}"), "},{")

	for _, v := range res1_2_1 {
		// fmt.Println(fmt.Sprintf("{%s}", v))
		var Test_1 T1
		json.Unmarshal([]byte(fmt.Sprintf("{%s}", v)), &Test_1)

		if Test_1.SPBM == "1070101070100000000" {
			fmt.Printf("润滑油结存为: %s\n", Test_1.SL)
		} else if Test_1.SPBM == "1070101070200000000" {
			fmt.Printf("润滑脂结存为: %s\n", Test_1.SL)
		}
	}
}
