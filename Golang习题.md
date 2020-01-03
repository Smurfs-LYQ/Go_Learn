1. 把字符串的IP地址"192.168.19.200"转换成整数
	```go
	package main

	import (
		"fmt"
		"math/big"
		"net"
	)

	// InetAtoN 将IP转换成整数
	func InetAtoN(ip string) int64 {
		ret := big.NewInt(0)
		ret.SetBytes(net.ParseIP(ip).To4())
		return ret.Int64()
	}

	// InetNtoA 将整数转换成IP
	func InetNtoA(ip int64) string {
		fmt.Println(ip)
		fmt.Printf("%T %v\n", ip>>24, ip>>24)
		return fmt.Sprintf("%d.%d.%d.%d",
			byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
	}

	func main() {
		ip := "192.168.19.200"
		ipInt := InetAtoN(ip)

		fmt.Printf("convert string ip [%s] to int: %d\n", ip, ipInt)
		fmt.Printf("convert int ip [%d] to string: %s\n", ipInt, InetNtoA(ipInt))
	}

	```

2. 变量d输出什么
	```go
	var a = []int{1,2,3,4,5,6,7}
	b := a[2:5]
	fmt.Println(c)		// [3, 4, 5]
	fmt.Println(len(c)) // 3
	fmt.Println(cap(c)) // 5

	c := b[:5]
	fmt.Println(c)		// [3, 4, 5, 6, 7]
	/*
		输出[3, 4, 5, 6, 7]的原因是b在切割完之后长度为3，所以b只输出[3，4，5]，但是的容量为5，其中包含的其实是[3, 4, 5, 6, 7]，但是只
	*/
	```

3. 查看下面的代码的输出结果
	```go
	var a = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	```

4. "hou do you do"中每个单词出现的次数
	```go
	str := "how do you do"
	var a = make(map[string]int, 3)
	for _, v := range strings.Split(str, " ") {
		_, ok := a[v]
		if ok {
			a[v]++
		} else {
			a[v] = 1
		}
	}
	fmt.Println(a)
	```

5. defer经典面试题
	```go
	package main

	import "fmt"

	func f1() int {
		x := 5
		defer func() {
			x++
		}()
		return x
	}

	func f2() (x int) {
		defer func() {
			x++
		}()
		return 5
	}

	func f3() (y int) {
		x := 5
		defer func() {
			x++
		}()
		return x
	}

	func f4() (x int) {
		defer func(x int) {
			x++
		}(x)
		return 5
	}

	func main() {
		fmt.Println(f1())
		fmt.Println(f2())
		fmt.Println(f3())
		fmt.Println(f4())
	}
	```

6. 为什么下面的代码输出 `真` ？
	```go
	func main() {
		f := func() bool {
			return false
		}

		switch f(); {
		case true:
			fmt.Println("真")
		case false:
			fmt.Println("假")
		}
	}
	```
	答: 
		因为 `f ()`后面的分号的将代码给分隔开了，`f()` 的返回值没有被接收，代码的效果等同于 `switch _ = f(); {` 、`switch {`，而 `switch {` 的判断条件为空的情况下就会默认将判断条件设置为布尔类型并且默认值为 `true`。

7. 下面代码会输出什么

	```go
	package main

	import "fmt"

	func T1() {
		defer func() { fmt.Println("print before recovery") }()
		defer func() { recover() }() // 如果recover不在你们函数内，那么panic报错就无法拦截，会在输出完defer的打印之后，执行panic
		defer func() { fmt.Println("print after recovery") }()
		panic("err")
	}

	func main() {
		T1()
	}
	```

	```txt
	答案
	print after recovery
	print before recovery
	```

8. 下面代码输出什么

	```go
	package main

	import "fmt"

	func calc(index string, a, b int) int {
		ret := a + b
		fmt.Println(index, a, b, ret)
		return ret
	}

	func main() {
		a := 1
		b := 2
		defer calc("1", a, calc("10", a, b))
		a = 0
		defer calc("2", a, calc("20", a, b))
		b = 1
	}
	```

	```txt
	答案
	10 1 2 3
	20 0 2 2
	2 0 2 2
	1 1 3 4

	执行流程
	calc("10", a, b)
	calc("20", a, b)
	defer calc("2", a, calc("20", a, b))
	defer calc("1", a, calc("10", a, b))
	```

9. 下面代码会不会触发异常

	```go
	package main

	import "runtime"

	import "fmt"

	func main() {
		runtime.GOMAXPROCS(1)
		for {
			int_chan := make(chan int, 1)
			string_chan := make(chan string, 1)
			int_chan <- 1
			string_chan <- "behe"
			select {
			case val := <-int_chan:
				fmt.Println(val + "out put")
			case val := <-string_chan:
				fmt.Println(val + "out put")
			}
		}
	}
	```

	```txt
	答案
		会，语法有问题，整型和字符串不能直接用 + 号拼接，
		如果语法问题修复了就不会报错了，因为chan是在for循环内定义的，每次循环都会重新定义
	```

10. 下面代码会输出什么

	```go
	package main

	import "fmt"

	type Parent struct{}

	func (p *Parent) MethodB() {
		fmt.Println("methodB from parent")
	}

	func (p *Parent) MethodA() {
		fmt.Println("methodA from parent")
		p.MethodB()
	}

	type Child struct {
		Parent
	}

	func (b *Child) MethodB() {
		fmt.Println("methodB from child")
	}

	func main() {
		child := Child{}
		child.MethodA()
	}
	```

	```txt
	答案
		methodA from parent
		methodB from parent

		注意methodA方法里是通过 p.MethodB 来调用的，也就是说他调用的是parent的MethodB
	```