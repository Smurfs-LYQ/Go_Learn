package demo

// One 编写一个回文检测函数，并为其编写单元测试和基准测试，根据测试的结果逐步对其进行优化。(回文: 一个字符串正序和逆序一样，如“油灯少灯油”)
func One(str string) bool {
	len := len(str)
	for i := 0; i < len; i++ {
		if str[i] != str[len-1] {
			return false
		}
		len--
	}
	return true
}
