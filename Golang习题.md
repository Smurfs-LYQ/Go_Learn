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

17. 在A和B处填入代码，使输出为foo

	```go
	package main

	type S struct {
		m string
	}

	func f() *S {
		return _ //A
	}

	func main() {
		p := _ //B
		println(p.m)
	}
	```

	```txt
	答案
		A: &S{"foo"} 或 &S{m:"foo"}
		B: f()
	```

18. 下面代码输出是什么，若想输出012，怎么改

	```go
	package main

	const N = 3

	func main() {
		m := make(map[int]*int)
		for i := 0; i < N; i++ {
			m[i] = &i
		}
		for _, v := range m {
			println(*v)
		}
	}
	```

	```txt
	答案
		333
		因为map创建的时候值设置的是指针类型，并且i执行最后一定要++到3的时候才会退出，而i的内存地址一直没有变过，所以结果输出的值都为3
	修改成输出012
		把代码中一些指针的操作都去掉，包括定义、取址和取值
	```

19. 代码输出什么？为什么？如何改会使得len(m)为10

	```go
	package main

	import "sync"

	const N = 10

	func main() {
		m := make(map[int]int)

		wg := &sync.WaitGroup{}
		mu := &sync.Mutex{}

		wg.Add(N)
		for i := 0; i < N; i++ {
			go func() {
				defer wg.Done()
				mu.Lock()
				m[i] = i
				mu.Unlock()
			}()
		}
		wg.Wait()
		println(len(m))
	}
	```

	```txt
	答案
		
		数据是不固定的，不会超过10不会小于1，因为这段代码是通过goroutine调用的一个匿名函数，而因为这个匿名函数调用了外部环境的变量，所以这个匿名函数也是一个闭包，所有的goroutine都会因为闭包共享同样的变量。原因是因为它是循环使用goroutine执行的，但是调用i是直接通过for循环调用的，并没有将参数i传入到匿名函数中，所以很有可能goroutine执行到需要调用i的时候，i已经循环过去好几个了。
	将i以参数的形式传入到匿名函数中长度就为10了
	```

20. 请描述golang语言的初始化顺序: 包，全局变量，全局常量，init函数、main函数

	```txt
	答案
		1. 包
		2. 全局常量
		3. 全局变量
		4. init函数
		5. main函数
	```

21. 定义一个全局字符串变量，下面哪个是对的

	```go
	1. var str string
	2. str := ""
	3. str = ""
	4. var str = ""
	```

	```txt
	答案
		1和4
	```

22. 描述下面代码输出

	```go
	package main

	import "fmt"

	type S1 struct {
	}

	func (s1 S1) f() {
		fmt.Println("S1.f()")
	}

	func (s1 S1) g() {
		fmt.Println("S1.g()")
	}

	type S2 struct {
		S1
	}

	func (s2 S2) f() {
		fmt.Println("S2.f()")
	}

	type I interface {
		f()
	}

	func printType(i I) {
		if s1, ok := i.(S1); ok {
			s1.f()
			s1.g()
		}
		if s2, ok := i.(S2); ok {
			s2.f()
			s2.g()
		}
	}

	func main() {
		printType(S1{})
		printType(S2{})
	}
	```

	```txt
	答案
		S1.f()
		S1.g()
		S2.f()
		S1.g()
	描述
		s1, ok := i.(S1) 这个地方是类型断言，进入对于的if之后调用f()和g()方法，s1好理解，因为他自己f和g两个方法都有，s2自身之声明了一个f方法，自己有的优先调用自己的元素，g方法s2没有，所以调用了嵌套的s1的f方法
	```

23. 下面代码有什么问题，怎么修改

	```go
	package main

	import (
		"math/rand"
		"sync"
	)

	const N = 10

	func main() {
		m := make(map[int]int)
		wg := &sync.WaitGroup{}

		wg.Add(N)
		for i := 0; i < N; i++ {
			go func() {
				defer wg.Done()
				m[rand.Int()] = rand.Int()
			}()
		}
		wg.Wait()
		println(len(m))
	}
	```

	```txt
	答案
		多个goroutine同时操作一个map会出现并发问题，只需要在操作map的时候加上一个互斥锁就可以了
		并且随机数也可能出现覆盖的问题
	```

24. 描述make和new的区别

	```txt
	new为指定的类型分配一片内存，初始化为0并且返回类型为指定类型指针的内存地址：这种方法 **返回一个指定类型，值为0的地址的指针**，它适用于值类型如数组和结构体；它相当于&T{}
	make**返回一个指定类型的初始值**，他使用与3中内建的引用类型：slice、map和channel
	new函数分配内存，make函数初始化

	1、new可以是任意类型，返回的是指针：*T，只分配内存，不初始化内存，只是将其置零
	2、make只能用于map，slice，chan，返回一个初始化的(而不是置零)，类型为T的值
	```

25. 下面代码输出什么，如何让输出为true

	```go
	package main

	import (
		"fmt"
	)

	type S struct {
		a, b, c string
	}

	func main() {
		x := interface{} (&S{"a", "b", "c"})
		y := interface{} (&S{"a", "b", "c"})
		fmt.Println(x == y)
	}
	```

	```txt
	答案
		默认输出false，因为程序拿的是地址，两个地址肯定是不相同的
		如果想输出true把取址符删掉就可以了
	```

26. 下面代码的问题是什么，如何修改

	```go
	package main

	type S struct {
		name string
	}

	func main() {
		m := map[string]S{"x": S{"one"}}
		m["x"].name = "two"
	}
	```

	```txt
	答案
		map里结构体无法直接寻址，必须取地
		将 map[string]S 改成 map[string]&S
	```

27. 修改代码，使的status输出为200

	```go
	package main

	import (
		"encoding/json"
		"fmt"
	)

	type Result struct {
		status int
	}

	func main() {
		var data = []byte(`{"status":200}`)
		result := &Result{}
		if err := json.Unmarshal(data, &result); err != nil {
			fmt.Println("error", err)
			return
		}
		fmt.Printf("result=%+v", result)
	}
	```

	```txt
	答案
		这个题的问题出在Result结构体的status字段首字母是小写的(不对外开放)，这导致json包在调用这个结构体的时候没法找到status这个字段，所以没法进行反序列化。
	```

28. 描述golang中的stack和heap的区别，分别在什么情况下会分配到stack？又在何时分配到heap中

	```txt
	区别：
    	1. heap是堆，stack是栈。
    	2. 栈(stack) : 由编译器自动分配和释放内存，存变量名、函数名等各种名
    	3. 堆(heap) : 在C里有程序员分配和释放内存，go是自动的，存栈中变量的值
    	4. stack空间有限，heap的空间是很大的自由区。
	分别在什么情况下会分配到stack？又在何时分配到heap中
		比如说 a := 3
		a 在stack里，3 在heap里
	```

29. 