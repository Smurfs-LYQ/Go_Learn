package main

//import "Go_learn/old/test/one"

func funcMui(x, y int) (int, error) {
	sum := x + y
	return sum, nil
}

func main() {
	//one.Do()
	println(funcMui(1, 2))
}
