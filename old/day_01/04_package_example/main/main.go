package main

import (
	"Go_learn/old/day_01/04_package_example/calc" //引入项目中的calc包
	"fmt"
)

func main() {
	add := calc.Add(1, 2)
	del := calc.Del(2, 1)

	fmt.Println(add)
	fmt.Println(del)
}
