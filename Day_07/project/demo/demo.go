package demo

import "unicode"

// One 编写一个回文检测函数，并为其编写单元测试和基准测试，根据测试的结果逐步对其进行优化。(回文: 一个字符串正序和逆序一样，如“油灯少灯油”)
// 优化前不可检测中文
func One(str string) bool {
	lenStr := len(str)
	for i := 0; i < lenStr; i++ {
		if str[i] != str[lenStr-1] {
			return false
		}
		lenStr--
	}
	return true
}

/*
// One执行结果
=== RUN   TestOne
=== RUN   TestOne/1
=== RUN   TestOne/2
=== RUN   TestOne/3
--- FAIL: TestOne (0.00s)
    --- PASS: TestOne/1 (0.00s)
    --- PASS: TestOne/2 (0.00s)
    --- FAIL: TestOne/3 (0.00s)
        demo_test.go:17: 油灯少灯油不是回文
FAIL
exit status 1
FAIL    Go_Learn/Day_07/project/one/demo        0.004s
*/

// 优化后可以检测中文并且忽略字符和字母大小写
func Two(s string) bool {
	var str []rune // 一个rune对应一个UTF-8字符, 这样就可以检测中文了
	for _, v := range s {
		// 只保留字符串中的字符串，剔除字符串中的特殊符号(标点符号等)
		if unicode.IsLetter(v) {
			// 忽略字符大小写并将其加入到str列表中
			str = append(str, unicode.ToLower(v))
		}
	}

	lenStr := len(str)
	for i := 0; i < lenStr/2; i++ {
		if str[i] != str[lenStr-i-1] {
			return false
		}
	}
	return true
}

/*
// Two执行结果
=== RUN   TestOne
=== RUN   TestOne/3
=== RUN   TestOne/1
=== RUN   TestOne/2
--- PASS: TestOne (0.00s)
    --- PASS: TestOne/3 (0.00s)
    --- PASS: TestOne/1 (0.00s)
    --- PASS: TestOne/2 (0.00s)
goos: darwin
goarch: amd64
pkg: Go_Learn/Day_07/project/one/demo
BenchmarkOne-16         52604770                21.7 ns/op             0 B/op          0 allocs/op
PASS
ok      Go_Learn/Day_07/project/one/demo        1.172s
*/
