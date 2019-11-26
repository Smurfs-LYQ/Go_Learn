package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func one() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go one()
	wg.Wait()
}
