package main

import "fmt"

type a struct {
	name string
}

func main() {
	one := map[string]*a{}

	one["1"] = &a{"1"}

	// one["3"] = &a{"3"}

	// if ok := one["4"]; ok != nil {
	// 	fmt.Println(ok)
	// }
	aaa(one)
	fmt.Println(one)
}

func aaa(one map[string]*a) {
	one["2"] = &a{"2"}
}
