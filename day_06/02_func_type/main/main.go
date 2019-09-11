package main

func Add(a, b int) int {
	return a+b
}

func Sum(a, b int) int {
	return a-b
}

type Op func (int, int) int

func Do(f Op, a, b int) int {
	return f(a, b)
}

func main() {
	println(Do(Add, 1, 2))
	println(Do(Sum, 1, 2))
}
