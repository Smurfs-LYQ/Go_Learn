package demo

import "sync"

import "fmt"

import "time"

var wg_2 sync.WaitGroup

func do_two(ch <-chan bool) {
	defer wg_2.Done()
	for {
		select {
		case <-ch:
			return
		default:
			fmt.Println("waiting...")
			time.Sleep(time.Second)
		}
	}
}

// Channel_demo channel关闭goroutine
func Channel_demo() {
	var ch1 = make(chan bool, 1)
	wg_2.Add(1)
	go do_two(ch1)

	time.Sleep(time.Second * 5)
	ch1 <- true
	wg_2.Wait()
	close(ch1)
	fmt.Println("over")
}
