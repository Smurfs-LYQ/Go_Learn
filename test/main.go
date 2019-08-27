package main

import "fmt"

type T1 struct {
	Name string
}

func main() {
	var one *T1 = T1{
		Name "asdf",
	}

	fmt.Println()
}
