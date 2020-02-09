#### <center>Day_12</center>

1. gob序列化示例
2. Context示例
3. Context - 客户端超时取消示例

#### <center>笔记</center>

1. > gob序列化示例

    - 标准库gob是golang提供的 "私有" 的编码方式，它的效率会比json、xml等更高，特别适合在Go语言程序间传递数据。

        ```go
        func GobDemo() {
            var s1 = model.S{
                Data: make(map[string]interface{}, 8),
            }
            s1.Data["count"] = 1

            // encode 编码
            buf := new(bytes.Buffer)   // 创建一个字节类型的缓冲区
            enc := gob.NewEncoder(buf) // 创建一个新的编码器对象
            err := enc.Encode(s1.Data) // 开始编码 编码后的数据会放在"创建编码器对象"时传入的缓冲区
            if err != nil {
                fmt.Println("gob encode failed, err:", err)
                return
            }

            b := buf.Bytes() // 拿到编码后的字节数据
            fmt.Println(b)

            var s2 = model.S{
                Data: make(map[string]interface{}, 8),
            }
            // decode 解码
            dec := gob.NewDecoder(bytes.NewBuffer(b)) // 创建一个新的解码器对象 并传入一个字节类型的缓冲区(里边保存需要解码的字节数据)
            err = dec.Decode(&s2.Data)                // 开始解码
            if err != nil {
                fmt.Println("gob decode failed, err:", err)
                return
            }
            fmt.Println(s2.Data)
            for _, v := range s2.Data {
                fmt.Printf("value: %v, type: %T\n", v, v)
            }
        }
        ```

2. > 什么是context

    - Go1.7加入了一个新的标准库 `context`, 它定义了 `Context` 类型，专门用来简化对于处理单个请求的多个goroutine之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。

    - 对服务器传入请求应该创建上下文，而对服务器的传出调用应该接受上下文。它们之间的函数调用链必须传递上下文，或者可以使用 `WithCancel`、`WithDeadline`、`WithTimeout` 或 `WithValue` 创建的派生上下文。当一个上下文被取消时，它派生的所有上下文也被取消。

3. > context - Context接口

    - `context.Context` 是一个接口，该接口定义了四个需要实现的方法。具体签名如下:

        ```go
        type Context interface {
            Deadline() (deadline time.Time, ok bool)
            Done() <-chan struct{}
            Err() error
            Value(key interface{}) interface{}
        }
        ```

        - `Deadline`方法需要返回当前`Context`被取消的时间，也就是完成工作的截止时间（deadline）；
        - `Done`方法需要返回一个`Channel`，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用`Done`方法会返回同一个Channel
        - `Err`方法会返回当前`Context`结束的原因，它只会在`Done`返回的Channel被关闭时才会返回非空的值
            - 如果当前`Context`被取消就会返回`Canceled`错误
            - 如果当前`Context`超时就会返回`DeadlineExceeded`错误
        - `Value`方法会从`Context`中返回键对应的值，对于同一个上下文来说，多次调用`Value`并传入相同的`Key`会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据

4. > context - Background()和TODO()

    - Go内置两个函数：`Background()` 和 `TODO()`，这两个函数分别返回一个实现了 `Context` 接口的 `background` 和 `todo`。我们代码中最开始都是以这两个内置的上下文对象作为最顶层的 `partent context`，衍生出更多的子上下文对象。

    - `Background()` 主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。

    - `TODO()`，它目前还不知道具体的使用场景，如果我们不知道该使用什么Context的时候，可以使用这个。

    - `background` 和 `todo` 本质上都是 `emptyCtx` 结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context。

5. > context - With系列函数 - WithCancel

    - `WithCancel` 的函数签名如下:

        ```go
        func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
        ```

        `WithCancel` 返回带有新Done通道的父节点的副本。当调用返回的cancel函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论发生什么情况。

    - 取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用 `cancel`。

        ```go
        func gen(ctx context.Context) <-chan int {
            dst := make(chan int)
            n := 1
            go func() {
                for {
                    select {
                    case <-ctx.Done():
                        return // return结束该goroutine，防止泄露
                    case dst <- n:
                        n++
                    }
                }
            }()
            return dst
        }
        func main() {
            ctx, cancel := context.WithCancel(context.Background())
            defer cancel() // 当我们取完需要的整数后调用cancel

            for n := range gen(ctx) {
                fmt.Println(n)
                if n == 5 {
                    break
                }
            }
        }
        ```

    - 上面的示例代码中，`gen` 函数在单独的goroutine中生成整数并将它们发送到返回的通道。gen的调用者在使用生成的整数之后需要取消上下文，以免 `gen` 启动的内部goroutine发生泄漏。

6. > context - With系列函数 - WithDeadline

    - `WithDeadline` 的函数签名如下:

        ```go
        func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
        ```

        返回父上下文的副本，并将deadline调整为不迟于d。如果父上下文的deadline已经早于d，则WithDeadline(parent,d)在语义上等同于父上下文。当截止日过期时，当调用返回的cancel函数时，或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。

    - 取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。

        ```go
        func main() {
            d := time.Now().Add(50 * time.Millisecond)
            ctx, cancel := context.WithDeadline(context.Background(), d)

            // 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
            // 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
            defer cancel()

            select {
            case <-time.After(1 * time.Second): // 定时器，到达预设的时间之后会返回当时的时间。
                fmt.Println("overslept")
            case <-ctx.Done(): // 到达预设的时间之后，ctx.Done()函数会接收到关闭信号
                fmt.Println(ctx.Err())
            }
        }
        ```

    - 上面的代码中，定义了一个50毫秒之后过期的deadline，然后我们调用 `context.WithDeadline(context.Background(), d)` 得到一个上下文(ctx) 和 一个取消函数(cancel)，然后使用一个select让主程序陷入等待: 等待1秒后打印 `overslept` 退出或者等待ctx过期后退出。因为ctx50毫秒后就过期，所以 `ctx.Done()` 会先接收到值，上面的代码会打印ctx.Err()取消原因。

7. > context - With系列函数 - WithTimeout

    - `WithTimeout` 的函数签名如下:

        ```go
        func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
        ```

        `WithTimeout` 返回 `WithDeadline(parent, time.Now().Add(timeout))`。

    - 取消此上下文将释放与其相关的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel，通常用于数据库或者网络连接的超时控制。具体示例如下:

        ```go
        var wg sync.WaitGroup

        func worker(ctx context.Context) {
        LOOP:
            for {
                fmt.Println("db connecting ...")
                time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
                select {
                case <-ctx.Done(): // 50毫秒后自动调用
                    break LOOP
                default:
                }
            }
            fmt.Println("worker done!")
            wg.Done()
        }

        func main() {
            // 设置一个50毫秒的超时
            ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
            wg.Add(1)
            go worker(ctx)
            time.Sleep(time.Second * 5)
            cancel() // 通知子goroutine结束
            wg.Wait()
            fmt.Println("over")
        }
        ```

8. > context - With系列函数 - WithValue

    - `WithValue` 函数能够将请求作用域的数据与Context对象建立关系。声明如下:

        ```go
        func WithValue(parent Context, key, val interface{}) Context
        ```

        `WithValue` 返回父节点的副本，其中与key关联的值为val。

    - 仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数。

    - 所提供的键必须是可比较的，并且不应该是 `string` 类型或任何其他内置类型，以避免使用上下文在包之间发生冲突。`WithValue` 的用户应该为键定义自己的类型。为了避免在分配给interface{}时进行分配，上下文键通常具有具体类型。或者，导出的上下文关键变量的静态类型应该是指针或接口。

        ```go
        type TraceCode string

        var wg sync.WaitGroup

        func worker(ctx context.Context) {
            key := TraceCode("TRACE_CODE")
            traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
            if !ok {
                fmt.Println("invalid trace code")
            }
        LOOP:
            for {
                fmt.Printf("worker, trace code:%s\n", traceCode)
                time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
                select {
                case <-ctx.Done(): // 50毫秒后自动调用
                    break LOOP
                default:
                }
            }
            fmt.Println("worker done!")
            wg.Done()
        }

        func main() {
            // 设置一个50毫秒的超时
            ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
            // 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
            ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
            wg.Add(1)
            go worker(ctx)
            time.Sleep(time.Second * 5)
            cancel() // 通知子goroutine结束
            wg.Wait()
            fmt.Println("over")
        }
        ```

9. > 使用Context的注意事项

    - 推荐以参数的方式显示传递Context
    - 以Context作为参数的函数方法，应该把Context作为第一个参数
    - 给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO()
    - Context的Value相关方法应该传递请求域的必要数据，不应该用于传递可选参数
    - Context是线程安全的，可以放心的在多个goroutine中传递

10. > os.Args

    - 如果你只是简单的想要获取命令行参数，可以像下面的代码示例一样使用 `os.Args` 来获取命令行参数。

        ```go
        package main

        import (
            "fmt"
            "os"
        )

        //os.Args demo
        func main() {
            //os.Args是一个[]string
            if len(os.Args) > 0 {
                for index, arg := range os.Args {
                    fmt.Printf("args[%d]=%v\n", index, arg)
                }
            }
        }
        ```

    - 运行结果

        ```go
        $ go run args_demo.go a b c d
        args[0]=./args_demo.go
        args[1]=a
        args[2]=b
        args[3]=c
        args[4]=d
        ```

    - `os.Args` 是一个存储命令行参数的字符串切片，它的第一个元素是执行文件的名称。

11. > flag包基本使用

    - Go语言内置的 `flag` 包实现了命令行参数的解析，`flag` 包使得开发命令行工具更为简单。

    - flag参数类型

        flag 包支持的命令行参数有 `bool`、`int`、`int64`、`uint`、`uint64`、`float`、`float64`、`string`、`duration`。

        | flag参数     | 有效值                                                       |
        | ------------ | ------------------------------------------------------------ |
        | 字符串flag   | 合法字符串                                                   |
        | 整数flag     | 1234、0664、0x1234等类型，也可以是负数                       |
        | 浮点数flag   | 合法浮点数                                                   |
        | bool类型flag | 1，0，t，f，T，F，true，false，TRUE，FALSE，True，False。    |
        | 时间段flag   | 任何合法的时间段字符串。如”300ms”、”-1.5h”、”2h45m”。合法的单位有”ns”、”us” /“µs”、”ms”、”s”、”m”、”h”。 |

    - 定义命令行flag参数 - flag.Type()

        基本格式如下:

        `flag.Type(flag名, 默认值, 帮助信息)*Type` 例如我们要定义姓名、年龄、婚否三个命令行参数，我们可以按照如下方式定义:

        ```go
        name := flag.String("name", "张三", "姓名")
        age := flag.Int("age", 18, "年龄")
        married := flag.Bool("married", false, "婚否")
        delay := flag.Duration("d", 0, "时间间隔")
        ```

        需要注意的是，此时 `name`、`age`、`married`、`delay` 均为对应类型的指针。

    - 定义命令行flag参数 - flag.TypeVar()

        基本格式如下: `flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)` 例如我们要定义姓名、年龄、婚否三个命令行参数，我们可以按如下方式定义:

        ```go
        var (
            name    string
            age     int
            married bool
            delay   time.Duration
        )
        flag.StringVar(&name, "name", "Smurfs", "姓名")
        flag.IntVar(&age, "age", 18, "年龄")
        flag.BoolVar(&married, "married", false, "婚否")
        flag.DurationVar(&delay, "delay", 0, "时间间隔")
        ```

    - flag.Parse()

        通过以上两种方法定义好命令行flag参数后，需要通过调用 `flag.Parse()` 来对命令行参数进行解析。

        支持的命令行参数格式有以下几种:

        - -flag xxx （使用空格，一个-符号）
        - --flag xxx （使用空格，两个-符号）
        - -flag=xxx （使用等号，一个-符号）
        - --flag=xxx （使用等号，两个-符号）

        其中，布尔类型的参数必须使用等号的方式指定。

        Flag解析在第一个非flag参数 (单个 "-" 不是flag参数) 之前停止，或者在终止符 "-" 之后停止。

    - flag其他函数

        ```go
        flag.Args()  // 返回命令行参数后的其他参数，以 []string 类型
        flag.NArg()  // 返回命令行参数后的其他参数个数
        flag.NFlog() // 返回使用的命令行参数个数
        ```

12. > Go性能优化

    - Go语言项目中的性能优化主要有以下几个方面:

        - CPU profile：报告程序的 CPU 使用情况，按照一定频率去采集应用程序在 CPU 和寄存器上面的数据
        - Memory Profile（Heap Profile）：报告程序的内存使用情况
        - Block Profiling：报告 goroutines 不在运行状态的情况，可以用来分析和查找死锁等性能瓶颈
        - Goroutine Profiling：报告 goroutines 的使用情况，有哪些 goroutine，它们的调用关系是怎样的

    - **采集性能数据**

        - Go语言内置了获取程序的运行数据的工具，包括以下两个标准库:

            - `runtime/pprof` : 采集工具型应用运行数据进行分析
            - `net/http/pprof` : 采集服务型应用运行时数据进行分析

        - pprof开启后，每隔10ms就会收集下当前的堆栈信息，获取个个函数占用的CPU以及内存资源；最后通过对这些采样数据进行分析，形成一个性能分析报告。

        - 注意: 我们只应该在性能测试的时候才在代码中引入pprof。

    - **工具型应用**

        - 如果你的应用程序是运行一段时间就结束退出类型。那么最好的办法是在应用退出的时候把 profiling 的报告保存到文件中，进行分析。对于这种情况，可以使用 `runtime/pprof` 库。

        - CPU性能分析

            开启CPU性能分析

            ```go
            pprof.StartCPUProfile(w io.Writer)
            ```

            停止CPU性能分析

            ```go
            pprof.StopCPUProfile()
            ```

            应用执行结束后，就会生成一个文件，保存了我们的 CPU profiling 数据。得到采样数据之后，使用 `go tool pprof` 工具进行CPU性能分析。

        - 内存性能分析

            记录程序的堆栈信息

            ```go
            pprof.WriteHeapProfile(w io.Writer)
            ```

            得到采样数据之后，使用 `go tool pprof` 工具进行内存性能分析。

            `go tool pprof` 默认是使用 `-inuse_space` 进行统计，还可以使用 `-inuse-objects` 查看分配对象的数量。

    - **服务型应用**

        - 如果你的应用程序是一直运行的，比如web应用，那么可以使用 `net/http/pprof` 库，它能够提供HTTP服务进行分析。

        - 如果使用了默认的 `http.DefaultServeMux` (通常是代码直接使用 http.ListenAndServe("0.0.0.0:8000", nil) )，只需要在你的web server端代码中导入 `net/http/pprof`

            如果你使用自定义的Mux，则需要手动注册一下路由规则:

            ```go
            r.HandleFunc("/debug/pprof/", pprof.Index)
            r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
            r.HandleFunc("/debug/pprof/profile", pprof.Profile)
            r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
            r.HandleFunc("/debug/pprof/trace", pprof.Trace)
            ```

        - 如果使用的是 `gin` 框架，那么推荐使用 `"github.com/DeanThompson/ginpprof"`。

        - 不管哪种方式，你的 HTTP 服务都会多出 `/debuf/pprof` endpoint，访问它会得到一下数据分析的信息，这个路径下还有几个子页面：

            /debug/pprof/profile：访问这个链接会自动进行 CPU profiling，持续 30s，并生成一个文件供下载

            /debug/pprof/heap： Memory Profiling 的路径，访问这个链接会得到一个内存 Profiling 结果的文件

            /debug/pprof/block：block Profiling 的路径

            /debug/pprof/goroutines：运行的 goroutines 列表，以及调用关系

    - **go tool pprof命令**

        - 不管是工具型应用还是服务型应用，我们使用相应的pprof库获取数据之后，下一步都要对这些数据进行分析，我们可以使用 `go tool pprof` 命令行工具。

        - `go tool pprof` 最简单的使用方式为:

            ```go
            go tool pprof [binart] [source]
            ```

            其中:

            ```txt
            binary: 是应用的二进制文件，用来解析各种符号。
            source: 表示 profile 数据的来源，可以是本地的文件，也可以是 http 地址。
            ```

        - 注意事项: 

            获取 Profiling 数据是动态的，要想获得有效的数据，请保证应用处于较大的负载 (比如正在生成中运行的服务，或者通过其他工具模拟访问压力)。否则如果应用处于空闲状态，得到的结果可能没有任何意义。