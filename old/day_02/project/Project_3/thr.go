package Project_3

/*
Project3
对于一个数n，求n的阶乘之和，即：1 * 2 * 3...
*/
func Project3(one int) int {
	num := 1
	res := 0
	for i := 1; i <= one; i++ {
		num = num * i
		res += num
	}
	return res
}
