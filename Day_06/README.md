#### <center>Day_06</center>

1. goroutine

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
5. 