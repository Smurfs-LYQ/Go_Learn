package goroute

func Add(a, b int, c chan int) {
	res := a + b
	c <- res // 将返回值res放入管道C中
}
