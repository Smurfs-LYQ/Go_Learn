package Project_2

import "fmt"

/*
Project2
打印出100-999中所有的“水仙花数”，所谓“水仙花数”是指一个三位数，其各位数字立方等于该数本身。例如：153是一个“水仙花数”，因为153=1的三次方+5的三次方+3的三次方
*/

func isNumber(n int) bool {
	var i, j, k int
	i = n % 10         // 实例 151 % 10 = 1, 因为前面的150可以取余 最后剩下1 因此就获取到了最后一位
	j = (n / 10) % 10  // 实例 (151 / 10) % 10 = 5, 前面151除10等于15.1 因为是整形自动去除小数点后的数，所以等于15，取余10 余5
	k = (n / 100) % 10 // 实例 (151 / 100) % 10 = 5, 前面151除10等于1.51 因为是整形自动去除小数点后的数，所以等于1，取余10 取不开 余1

	sum := i*i*i + j*j*j + k*k*k
	return sum == n
}

func Project1() {
	var n int // 开始值
	var m int // 结束值

	fmt.Scanf("%d,%d", &n, &m)

	for i := n; i < m; i++ {
		if isNumber(i) == true {
			fmt.Println(i)
		}
	}
}
