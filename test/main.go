package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	fmt.Println(rand.Intn(100))

	fmt.Println(rand.Float32())
}
