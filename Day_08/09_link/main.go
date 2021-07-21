package main

import "fmt"

type Moto struct {
	name string
}

func (m Moto) start() Moto {
	fmt.Printf("%s 点火\n", m.name)
	return m
}

func (m Moto) stop() Moto {
	fmt.Printf("%s 熄火\n", m.name)
	return m
}

func main() {
	Ducati := Moto{"杜卡迪V4R"}
	// 链式操作
	Ducati.start().stop()
}
