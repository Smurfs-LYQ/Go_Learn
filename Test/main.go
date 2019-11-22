package main

import "fmt"

func main() {
	for {
		go fmt.Println(1)
	}
}
