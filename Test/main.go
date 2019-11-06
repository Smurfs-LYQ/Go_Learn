package main

import (
	"fmt"
	"os"
)

func main() {
	res, err := os.Stat("./Test")
	fmt.Println(res.Size()/1024/1024, err)
}
