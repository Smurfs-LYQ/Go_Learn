package main

import "fmt"

func add(a, b float64) float64 {
	return float64(a) + float64(b)
}

func main() {
	fmt.Println(add(1, 1.1))
}
