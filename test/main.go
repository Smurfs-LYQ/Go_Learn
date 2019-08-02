package main

import "fmt"

func main() {
	a := map[int]string{
		1: "one",
		2: "two",
		3: "thr",
	}

	for key, val := range a {
		fmt.Println(key, val)
	}
}
