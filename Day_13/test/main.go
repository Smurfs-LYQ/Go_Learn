package main

import "fmt"

func main() {
	one := map[string]int{}
	one["1"] = 1
	one["2"] = 2
	one["3"] = 3

	for k, _ := range one {
		delete(one, k)
	}

	fmt.Println(one)
}
