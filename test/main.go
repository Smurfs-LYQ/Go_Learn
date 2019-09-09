package main

import (
<<<<<<< HEAD
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[1:])
=======
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for i := 0; i < 5; i++ {
		input.Scan()           // 让用户进行标准输入
		counts[input.Text()]++ // input.Text() 可以获取到刚刚用户输入的内容
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
>>>>>>> 3f65157f005765a2ad5829eb285fff4eaa1b4fa4
}
