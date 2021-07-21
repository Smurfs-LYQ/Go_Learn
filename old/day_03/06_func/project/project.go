package project

/*
	写一个函数Project_1, 支持1个或多个strings相拼接, 并返回结果
*/
func Project_1(one ...string) (res string) {
	for _, v := range one {
		res = res + v
	}
	return res
}
