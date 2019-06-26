package main

import (
	"Go_learn/day_01/04_package_example/calc"
	"fmt"
)

func main() {
	add := calc.Add(1, 2)
	del := calc.Del(2, 1)

	fmt.Println(add)
	fmt.Println(del)
}
