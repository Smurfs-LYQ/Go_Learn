package main

import "fmt"

type T1 struct {
	Name string
	Age  int
}

func (obj T1) String() string {
	str := fmt.Sprintf("name=[%s] age=[%d]", obj.Name, obj.Age)
	return str
}

func main() {
	var one T1
	one.Name = "test_1"
	one.Age = 18
	fmt.Printf("%s\n", one)
}
