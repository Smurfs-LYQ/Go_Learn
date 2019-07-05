package Project_3

/*
Project3
对于一个数n，求n的阶乘之和，即：1 * 2 * 3...
*/
func Project3(one int) int {
	res := 1
	for i := 1; i <= one; i++ {
		res = res * i
	}
	return res
}
