package thr

import "fmt"

// One 打印200~1000之间的质数(素数)
func One() {
	// 在外侧的循环处加一个退出用的Label标签
exitTag:
	for i := 200; i <= 1000; i++ {
		for o := 2; o < i; o++ {
			// 判断i被o取余是否为0，为0就表示当前i可以被当前o整除
			if i%o == 0 {
				// 如果取余为0，跳出当前外侧循环，进行下一次循环
				continue exitTag
			}
		}
		// 如果执行完for循环没有出现被取余为0的情况，则打印变量i
		fmt.Println(i)
	}
}
