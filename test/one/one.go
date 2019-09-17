package one

import (
	"fmt"
	"time"
)

func Do() {
	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
	}()

	go da()
	go db()
	time.Sleep(3 * time.Second)
}

func da() {
	panic("panic da")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func db() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
