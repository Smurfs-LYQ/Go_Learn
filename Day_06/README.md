#### <center>Day_06</center>

1. goroutine
2. goroutine与线程
3. channel
4. channel-无缓冲通道和有缓冲通道
5. channel-取值时判断通道是否关闭
6. channel-select多路复用
7. channel-单向通道
8. 并发控制与锁-互斥锁
9. 并发控制与锁-读写锁

#### <center>笔记</center>
1. > 并发和并行的区别
	- 并发: 同一 `时间` 同时在做多个事情 (你用微信和两个人聊天)
	- 并行: 同一 `时刻` 同时在做多个事情 (你和你的朋友都在用微信和朋友聊天)
	- Go语言的并发通过 `goroutine` 实现。`goroutine` 类似于线程，属于用户态的线程，也就是协程，我们可以根据需要创建成千上万个 `goroutine` 并发工作。`goroutine` 是由Go语言的运行时调度完成，而线程是由操作系统调度完成。
	- Go语言还提供 `channel` 在多个 `goroutine` 间进行通信。`goroutine` 和 `channel` 是Go语言秉承着CSP(Communicating Sequential Process)并发模式的重要实现基础。
2. > 进程、线程、协程
	- 进程: 一个程序启动之后就创建了一个进程
	- 线程: 操作系统调度的最小单位
	- 协程: 用户态的线程
3. > 使用goroutine
    - Go程序中使用 `go` 关键词为一个函数创建一个 `goroutine`。一个函数可以被创建多个 `goroutine`，一个 `goroutine` 必定都应一个函数。
	- 启动单个goroutine
		- 启动 `goroutine` 的方式非常简单，只需要在调用的函数(普通函数和匿名函数)前面加上一个 **`go`** 关键字。
		- 实例
			```
			func hello() {
				fmt.Println("hello Golang")
			}

			func main() {
				defer fmt.Println("程序结束")
				/*
					1. 创建一个goroutine
					2. 在新的goroutine中执行hello函数
				*/
				go hello()

				fmt.Println("hello world")
				time.Sleep(time.Second) // 因为创建goroutine的时候程序就已经执行完了，所以需要加上这个sleep函数让程序等一会goroutine
			}
			```
		- 在程序启动时，Go程序就会为 `main` 函数创建一个默认的 `goroutine`。当main()函数返回的时候该 `goroutine` 就结束了，所有在 `main()` 函数中启动的 `goroutine` 会一同结束，`main` 函数就像是权利的游戏中的夜王，其他的 `goroutine` 都是异鬼，夜王意思它转化的那些异鬼就全部GG了。所以我们要想办法让main函数等一等hello函数，最简单粗暴的方式就是 `Sleep`。
		- 直接结果会先打印 `hello world`，然后紧接着打印 `hello Golang`，这是因为创建 `goroutine`的时候需要花费一些时间，而此时main函数是继续向下执行的。
    - **sync.WaitGroup**
        - 在代码中生硬的使用 `time.Sleep` 肯定是不合适的，Go语言中可以使用 `sync.WaitGroup` 来实现并发任务的同步，它有一下几个方法。
			|   方法名			|  功能  |
			|   ---   		   |  ---  |
			| (wg * WaitGroup) Add(i int) | 计数器 + i |
			| (wg * WaitGroup) Done() | 计数器 - 1 |
			| (wg * WaitGroup) Wait() | Wait方法阻塞直到WaitGroup计数器减为0 |
        - 实例
			```
			// 实例化sync包中的WaitGroup结构体(它里面有一个计数器)
			var wg sync.WaitGroup

			func hello(i int) {
				defer wg.Done() // 计数器-1, wg.Add()中的参数-1

				fmt.Println("hello Golang", i)
			}

			func main() {
				defer fmt.Println("程序结束")

				wg.Add(10) // 创建一个等待标签, 参数为: 需要等待的次数
				for i := 0; i < 10; i++ {
					go hello(i)
				}

				fmt.Println("hello world")

				// 等hello执行完(执行hello函数的那个goroutine执行完)
				wg.Wait() // 阻塞，一直等待所有的goroutine结束。当wg.Add()中的参数减成0的时候才会执行wg.Wait
			}
			```
        - 多次执行上面的代码，会发现每次打印的数字的顺序都不一致。这是因为10个 `goroutine` 是并发执行的，而 `goroutine` 的调度是随机的。
4. > goroutine与线程
	- Go语言中的操作系统线程和goroutine的关系:
    	- 一个操作系统线程对应用户态多个goroutine
    	- go程序可以同时使用多个操作系统线程
    	- goroutine和OS线程是多对多的关系，即m:n
	- 可增长的栈
    	- OS线程(操作系统线程)一般都是固定的栈内存(通常为2MB)，一个 `Goroutine` 的栈在其生命周期开始时只有很小的栈(典型情况下2KB)，`goroutine` 的栈不是固定的，他可以按需增大和减小，`goroutine` 的栈大小限制可以达到1GB，虽然极少会用到这么大。所以在Go语言中一次创建十万左右的 `goroutine` 也是可以的。
	- goroutine调度
    	- OS线程是由OS内核来调度的，`goroutine` 则是由Go运行时(runtime)自己的调度器调度的，这个调度器使用一个称为m:n调度的技术(复用/调度m个goroutine到n个OS线程)。goroutine的调度不需要切换内核语境，所以调用一个goroutine比调度一个线程成本低很多。
	- runtime.GOMAXPROCS
    	- Go运行时的调度器使用 `GOMAXPROCS` 参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上(GOMAXPROCS是m:n调度中的n)。
    	- Go语言中可以通过 `runtime.GOMAXPROCS()` 函数设置当前程序并发时占用的CPU逻辑核心数。
    	- Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。
5. > channel
	```
		单纯的将函数并发执行是没有意义的，函数与函数间需要交换数据才能体现并发执行函数的意义。
		虽然可以使用共享内存进行数据交换，但是共享内存在不同的"goroutine"中容易发生竞态问题。为了保证数据交换的正确性，必须使用"互斥量"对内存进行加锁，这种做法势必造成性能问题。
		Go语言的并发模型是CSP，提倡通过通信共享内存而不是通过共享内存而实现通信。
		如果说"goroutine"是Go程序并发的执行体，"channel"就是它们之间的连接。"channel"是可以让一个"goroutine"发送特定值到另一个"goroutine"的通信机制。
		Go语言中的通道(channel)是一种特殊的类型。通道像一个传送带或这队列，总是遵循先入先出(First In First Out)的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
	```
	- **声明channel**
    	- 声明通道类型的格式如下:
			```
			var 变量 chan 元素类型
			```
    	- 举几个例子:
			```
			var ch1 chan int	// 声明一个传递整型的通道
			var ch2 chan bool	// 声明一个传递布尔型的通道
			var ch3 chan []int	// 声明一个传递int切片的通道
			```
	- **创建channel**
    	- 通道是引用类型，通道类型的零值是 `nil`。
    	- 声明通道后需要使用 `make` 函数初始化之后才能使用。
			```
			make(chan 元素类型, [缓冲大小])
			缓冲大小是可选的
			```
	- **channel操作**
    	- 通道有发送(send)、接收(receive)和关闭(close)三种操作。发送和接收都是用 `<-` 符号。
    	- 讲一个值发送到通道中
			```
			ch <- 10 // 把10发送到ch中
			```
    	- 从通道中接收值
			```
			res := <- ch // 从ch中接收值并赋值给变量x
			<- ch		 // 从ch中接收值，忽略结果
			```
    	- 我们通过调用内置的 `close` 函数来关闭通道
			```
			close(ch)
			```
    	- 关于关闭通道需要注意的事情是，只有在通知接收方 `goroutine` 所有的数据都发送完毕的时候才需要关闭通道。通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但是关闭通道不是必须的。
    	- 关闭通道有以下特点:
        	- 对一个关闭的通道再发送值会导致panic
        	- 关闭一个已经关闭的通道会导致panic
        	- 对一个关闭的通道进行接收会一直获取值，直到通道为空。
        	- 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
6. > 无缓冲的通道和有缓冲的通道
	- 无缓冲的通道又被称为阻塞的通道，只有在有人接收值得时候才能发送值。
	- 只要通道的容量大于零，那么该通道就是有缓冲区的通道，通道的容量表示通道中能存放元素的数量。
	- 我们可以使用内置的 `len()` 函数获取通道内元素的数量，使用 `cap()` 函数获取通道的容量。
7. > 取值时，判断通道是否关闭
	- 方法一
		```
		result, ok := <- 管道变量
		```
		- 使用这种方式取值，当管道关闭的时候 ok = false
    - 方法二
		```
		for result := ch1 {
			fmt.Println(result)
		}
		```
		- 这种写法会一直从ch1管道中取值，程序会自动做类似于方法一那种判断，直到**管道关闭才会结束循环**
8. > select多路复用
	- 在某些场景下我们需要同时从多个通道接收数据。为了应对这种场景，Go内置了 `select` 关键字，可以同时相应多个通道的操作。select的使用类似于switch语句，它有一些列case分支和一个默认的分支。每个case会对应一个通道的通信(接收或发送)过程。`select` 会一直等待，直到某个 `case` 的通信操作完成时，就会执行 `case` 分支对应的语句。
	- 实例:
		```
		select {
			case ch1 <- 1:
				...
			case res := <-ch1:
				...
			case <-ch2:
				...
			default:
				...
		}
		```
	- 使用 `select` 语句能提高代码的可读性。如果多个 `case` 同时满足，`select` 会随机选择一个。对于没有 `case` 的 `select{}` 会一直等待。
9. > 单向通道
	- 方法一，接收一个**既能发送也能接收**的通道 (非单向通道)
		```
		func T1(ch chan int) {

		}
		```
	- 方法二，接收一个只能发送的通道 (单向通道)
		```
		func T2(ch chan<- int) {

		}
		```
	- 方法三，接收一个只能接收的通道 (单向通道)
		```
		func T3(ch <-chan int) {

		}
		```
10. > 并发控制
	- 有时候Go代码中可能会存在多个 `goroutine` 同时操作一个资源(临界区)，这种情况会发生 `竞态问题` (数据竞态)。类比显示生活中的例子有十字路口被各个方向的汽车竞争等。举个例子。
		```
		var x int64
		var wg sync.WaitGroup

		func add() {
			for i := 0; i < 1000; i++ {
				x++
			}
			wg.Done()
		}

		func main() {
			wg.Add(2)
			go add()
			go add()
			wg.Wait()
			fmt.Println(x)
		}
		```
	- 上面代码中开启了两个 `goroutine` 去累加变量x的值，这两个 `goroutine` 在访问和修改 `x` 变量的时候回存在数据竞争，导致最后的结果与期待的不符
11. > 锁
	- 互斥锁
		- 互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个 `goroutine` 可以访问共享资源。Go语言中使用 `sync` 包的 `Mutex` 类型来实现互斥锁。使用互斥锁来修复上面的代码问题:
			```
			var x int64
			var wg sync.WaitGroup
			var lock sync.Mutex

			func add() {
				for i := 0; i < 1000; i++ {
					// 加锁
					lock.Lock()

					x++

					// 解锁
					lock.Unlock()
				}
				wg.Done()
			}

			func main() {
				wg.Add(2)
				go add()
				go add()
				wg.Wait()
				fmt.Println(x)
			}
			```
		- 使用互斥锁能够保证同一时间有且只有一个 `goroutine` 进入临界区，其它的 `goroutine` 则在等待锁；当互斥锁释放后，等待的 `goroutine` 才可以获取锁进入临界区，多个 `goroutine` 同时等待一个锁时，唤醒的策略是随机的。
  	- 读写互斥锁
    	- 互斥锁是完全互斥的，但是当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下读写锁是更好的一种选择。读写锁使用 `sync` 包中 `RWMutex` 类型。
    	- 读写锁分为两种: `读锁`和`写锁`。当一个 `goroutine` 获取读锁之后，其他的 `goroutine` 如果是获取读锁会继续获得锁，如果是获取写锁就会等待；当一个 `goroutine` 获取读写锁之后，其他的 `goroutine` 无论是获取读锁还是写锁都会等待。
    	- 读比写多的时候使用读写锁 能够提高性能
12. 