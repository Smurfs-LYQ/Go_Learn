package main

import "fmt"

func main() {
	panic(fmt.Errorf("错误提示%d\n", 123))
}
