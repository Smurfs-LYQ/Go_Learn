package main

import (
	"fmt"
	"math/rand"
	"time"
)

func run(ch chan int) {
	for {
		for i := range time.Tick(time.Second) {
			fmt.Println(i)
			ch <- rand.Intn(100)
		}
	}
}

func main() {
	ch := make(chan int, 10)
	go run(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
