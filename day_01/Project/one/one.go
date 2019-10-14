package one

// One 实现字符串倒序
func One(str1 string) string {
	runeStr1 := []rune(str1)
	/*
		// 方法一
		var runeStr2 []rune
		for i := len(runeStr1) - 1; i >= 0; i-- {
			runeStr2 = append(runeStr2, runeStr1[i])
		}
		return string(runeStr2)
	*/
	/*
		// 方法二
		var str string
		for i := len(runeStr1) - 1; i >= 0; i-- {
			str = str + string(runeStr1[i])
		}
		return str
	*/
	// 方法三
	length := len(runeStr1) - 1
	for i := 0; i < length/2; i++ {
		runeStr1[i], runeStr1[length-i] = runeStr1[length-i], runeStr1[i]
	}
	return string(runeStr1)
}
