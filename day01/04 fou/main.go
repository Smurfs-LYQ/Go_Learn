package main

// import "fmt"

func main() {
	var x, y int
	print(&x, "\n")
	print(&y, "\n")
	println("------------")
	println(&x == &x, &x == &y, &x == nil)
}