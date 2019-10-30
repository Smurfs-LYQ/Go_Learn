package main

import "fmt"

func f1() int { // 不指定返回值变量
	x := 5
	defer func() {
		x++
	}()
	return x // 1. 返回值=5  2. x++  3. RET ==> 5
}

func f2() (x int) { // 指定返回值变量
	defer func() {
		x++
	}()
	return 5 // 1. 返回值=x (x=5)  2. x++  3. RET ==> x (x=6)
}

func f3() (y int) { // 指定返回值变量
	x := 5
	defer func() {
		x++
	}()
	return x // 1. 返回值=y (y=x, x=5, y=5)  2. x++  3. RET ==> y (y=5)
}

func f4() (x int) { // 指定返回值变量
	defer func(x int) {
		x++
	}(x) // 这里的x是值拷贝
	return 5 // 1. 返回值=x (x=5)  2. x++(函数内部的x) 3. RET ==> x (x=5)
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
