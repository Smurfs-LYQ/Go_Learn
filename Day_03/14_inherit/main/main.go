package main

import "fmt"

// Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

// Cat 猫
type Cat struct {
	cay     string
	*Animal // 通过嵌套匿名结构体实现继承
}

func (c *Cat) miao() {
	fmt.Printf("%s会%s~\n", c.name, c.cay)
}

func main() {
	T1 := &Cat{
		cay: "喵喵喵",
		Animal: &Animal{ // 注意嵌套的是结构体指针，也可以不嵌套结构体指针
			name: "火锅",
		},
	}
	T1.move()
	T1.miao()
}
