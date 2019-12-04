package main

import (
	"fmt"
	"math/rand"
)

var T1ch chan *T1
var T2ch chan *T2

type T1 struct {
	id  int
	num int
}

type T2 struct {
	*T1
	sum int
}

// Producter 生产者
func Producter(ch chan *T1) {
	for i := 0; i < 100; i++ {
		ch <- &T1{
			i,
			rand.Intn(1000),
		}
	}
	close(ch)
}

// Consummer 消费者
func Consummer(T1ch chan *T1, T2ch chan *T2) {
	for res := range T1ch {
		T2ch <- &T2{
			res,
			calc(res.num),
		}
	}
	/*
		T1 := <-T1ch
		T2ch <- &T2{
			T1,
			calc(T1.num),
		}
	*/
}

func calc(num int) (sum int) {
	for num > 0 {
		sum += num % 10
		num = num / 10
	}

	return
}

func printResult(T2ch chan *T2) {
	for res := range T2ch {
		fmt.Printf("id: %d num: %d sum: %d\n", res.id, res.num, res.sum)
		if len(T2ch) == 0 {
			close(T2ch)
		}
	}
	/*
		for {
			select {
			case res := <-T2ch:
				fmt.Printf("id: %d num: %d sum: %d\n", res.id, res.num, res.sum)
			}
		}
	*/
}

func main() {
	T1ch = make(chan *T1, 100)
	T2ch = make(chan *T2, 100)

	go Producter(T1ch)
	for i := 0; i < 10; i++ {
		go Consummer(T1ch, T2ch)
	}

	printResult(T2ch)
}
