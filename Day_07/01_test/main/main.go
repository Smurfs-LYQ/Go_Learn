package main

import (
	"Go_Learn/Day_07/01_test/split"
	"fmt"
)

func main() {
	for _, v := range split.Split("1:2:3:4", ":") {
		fmt.Println(v)
	}
}
