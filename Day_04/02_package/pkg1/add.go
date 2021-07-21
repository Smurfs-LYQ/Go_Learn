package pkg1

var t1 = 100 // 外部不可调用

const T2 = 200 // 外部可以调用

func Add(x, y int) int { // 外部可以调用
	return x + y
}

func del(x, y int) int { // 外部不可调用
	return x - y
}
