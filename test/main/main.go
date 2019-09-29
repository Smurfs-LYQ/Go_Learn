package main

//import "Go_Learn/test/one"

func funcMui(x,y int)(sum int,error) {
	sum := x+y
	return sum,nil
}

func main() {
	//one.Do()
	println(funcMui(1,2))
}
