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

11. 下面的代码有什么问题

	```go
	package main

	import (
		"fmt"
		"strconv"
		"sync"
	)

	type UserAges struct {
		ages map[string]int
		sync.Mutex
	}

	func (u *UserAges) Add(name string, age int) {
		u.Lock()
		defer u.Unlock()
		u.ages[name] = age
	}

	func (u *UserAges) Get(name string) int {
		if age, ok := u.ages[name]; ok {
			return age
		}
		return -1
	}

	func main() {
		u := UserAges{ages: map[string]int{}}
		wg := sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				u.Add("u"+strconv.Itoa(i), i)
				wg.Done()
			}(i)
		}
		// time.Sleep(time.Second)
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				fmt.Println(u.Get("u" + strconv.Itoa(i)))
				wg.Done()
			}(i)
		}
		wg.Wait()
		fmt.Println("done")
	}
	```

	```txt
	答案
		问题出在两个for循环中间没有设置时间间隔，有可能出现还没设置就已经读了的报错，
	解决办法
		把互斥锁改成读写锁，在读的时候加上读锁
	```

12. 下面代码输出什么

	```go
	package main

	import (
		"encoding/json"
		"fmt"
		"reflect"
	)

	func main() {
		json_str := []byte(`{"age":1}`)
		var value map[string]interface{}
		json.Unmarshal(json_str, &value)
		age := value["age"]
		fmt.Println(reflect.TypeOf(age))
	}
	```

	```txt
	答案
		float64
	原因
		源码中认为所有的数字都是float64
	```

13. 下面代码输出什么

	```go
	package main

	import (
		"fmt"
		"sync"
	)

	func main() {
		wg := sync.WaitGroup{}
		wg.Add(200)

		for i := 0; i < 10; i++ {
			go func() {
				fmt.Println("i ", i)
				wg.Done()
			}()
		}

		for i := 0; i < 10; i++ {
			go func(i int) {
				fmt.Println("j ", i)
				wg.Done()
			}(i)
		}
		wg.Wait()
	}
	```

	```txt
	答案
		i的值不固定,但一般不会是0-9，j是0-9
	原因
		并发协程传参会乱的
		这道题循环中的协程想要传参的话必须从匿名函数那个位置传入，像第一个循环直接调用可能会因为for执行的速度过快导致协程中的i被覆盖掉
	```

14. 下面的Max函数有什么问题

	```go
	package main

	import "math"

	import "fmt"

	func max(a, b int64) int64 {
		// math.Max的作用是对比这两个数值哪个是最大的
		return int64(math.Max(float64(a), float64(b)))
	}

	func main() {
		fmt.Println(max(1, 2))
		fmt.Println(max(math.MaxInt64-2, math.MaxInt64-1), math.MaxInt64-1)
	}
	```

	```txt
	答案
		Max有可能会导致数字溢出问题，
	```

15. 下面的迭代方法有什么问题

	```go
	package main

	import (
		"fmt"
		"reflect"
		"sync"
	)

	type threadSafeSet struct {
		sync.RWMutex
		s []interface{}
	}

	func (set *threadSafeSet) Iter() <-chan interface{} {
		ch := make(chan interface{})

		go func() {
			set.RLock()
			for _, val := range set.s {
				ch <- val
			}
			close(ch)
			set.RUnlock()
		}()

		return ch
	}
	```

	```txt
	答案
		因为ch是无缓冲区的通道，而for循环又一直在往通道里面传值，所以这里会导致卡死
	```

16. 列举chan的几种用法

	```txt
	读写数据
	有无缓冲
	workpool限流模型
	定时器
	关闭通道
	```
