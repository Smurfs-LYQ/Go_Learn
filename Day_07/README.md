#### <center>Day_07</center>

1. 单元测试
2. 单元测试组
3. 子测试
4. 基准测试
5. 性能比较函数
6. 并行测试
7. Setup 与 TearDown
8. 子测试的Setup 与 TearDown
9. 示例函数
10. TCP编程-Socket编程
11. UDP编程-Socket编程

#### <center>笔记</center>
1. > 单元测试
    - 单元组件是什么
      
      - 单元组件可以使函数、结构体、方法和最终用户可能依赖的任意东西。总之我们需要确保这些组件是能够正常运行的。
      
    - 如何测试单元组件的功能
      
      - 创建单元测试。单元测试是一些利用各种方法测试单元组件的程序，它会将结果与预期输出进行比较
      
    - 单元测试要测什么
      
      - 如果是一个模块或包，我们应该测试包中任意可导出使用的单位。如果我们有一个可执行包，我们应该测试任何在包范围内可用的单位。
      
    - **go test**
      - Go语言中的测试依赖 `go test` 命令。编写测试代码和编写普通的Go代码过程是类似的。
      
      - `go test` 命令是一个按照一定约定和组织的测试代码的驱动程序。在包目录内，所有以 `_test.go` 为后缀名的源代码文件都是 `go test` 测试的一部分，不会被 `go build` 编译到最终的可执行文件中。
      
      - 在 `*_test.go` 文件中有三种类型的函数: 单元测试函数、基准测试函数和示例函数。
      
        | 类型     | 格式                   | 作用                           |
        | -------- | ---------------------- | ------------------------------ |
        | 测试函数 | 函数名前缀为 Test      | 测试程序的一些逻辑行为是否正确 |
        | 基准函数 | 函数名前缀为 Benchmark | 测试函数的性能                 |
        | 示例函数 | 函数名前缀为 Example   | 为文档提供示例文档             |
      
      - `go test` 命令会遍历所有的 `*_test.go` 文件中符合上述命名规则的函数，然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。
      
      - `go test` 命令的options
      
        - 运行单元测试并返回详细信息
      
          ```go
          -v
          ```
      
        - 只想运行单元测试文件中的指定函数
      
          ```go
          第一种写法: -run 正则表达式
          第二种写法: -run="正则表达式"
          ```
      
        - 
      
   - 测试函数
   
     - 每个测试函数必须导入 `testing` 包，测试函数的基本格式 (签名) 如下:
   
       ```go
       func TestName(t *testing.T) {
         // ...
       }
       ```
   
     - 测试函数的名字必须以 `Test` 开头，可选的后缀名必须以大写字母开头。示例：
   
       ```go
       func TestAdd(t *testing.T) { ... }
       func TestSum(t *testing.T) { ... }
       func TestLog(t *testing.T) { ... } 
       ```
   
     - 其中参数 `t` 用于报告测试失败和附加的日志信息。`testing.T` 的拥有的方法有很多，详情请翻阅官方文档。
   
2. > 测试组

   - 测试组就是将多个测试用例同时放入到一个测试函数中进行测试。

   - 实例:

     ```go
     // 将多个测试用例放到一起组成 测试组
     func TestSplitGroup(t *testing.T) {
     	// 定义一个存放测试数据的结构体
     	type test struct {
     		str  string   // 字符串
     		sep  string   // 切割字符
     		want []string // 期望得到的值
     	}
     
     	// 创建一个存放多个测试用例的map
     	var tests = map[string]test{
     		"normal": test{"1,2,3", ",", []string{"1", "2", "3"}},
     		"none":   test{"1:2:3", ":", []string{"1", "2", "3"}},
     		"multi":  test{"1:2:3", ":2:", []string{"1", "2", "3"}},
     	}
     
     	// 循环调用测试用例
     	for k, v := range tests {
     		res := Split(v.str, v.sep) // 将测试用例中的数据放入到测试的函数中
     		if !reflect.DeepEqual(res, v.want) {
     			t.Errorf("测试用例: %v失败, 期望找到: %v, 实际得到: %v\n", k, v.want, res)
     		}
     	}
     }
     ```

     

3. > 子测试

   - 测试用例比较多的时候，我们没办法一眼看出来具体是哪个测试用例失败了，这时我们可以向上一个示例一样直接将失败的测试用例打印出来，**或者使用Go1.7+中新增的`子测试`，我们可以按照如下方式使用 `t.Run` 执行子测试**

     示例

     ```go
     // 将多个测试用例放到一起组成 测试组
     func TestSplit(t *testing.T) {
       // 定义一个存放测试数据的结构体
       type test struct {
         str  string   // 字符串
         sep  string   // 切割字符
         want []string // 期望得到的值
       }
     
       // 创建一个存放多个测试用例的map
       var tests = map[string]test{
         "normal": test{"1,2,3", ",", []string{"1", "2", "3"}},
         "none":   test{"1:2:3", ":", []string{"1", "2", "3"}},
         "multi":  test{"1:2:3", ":2:", []string{"1", "2", "3"}},
       }
     
       // 循环调用测试用例
       for k, v := range tests {
         t.Run(k, func(t *testing.T) { // 使用t.Run()执行子测试 参数: 测试用例名字，测试用例执行的函数(函数变量或匿名函数)
           res := Split(v.str, v.sep)  // 将测试用例中的数据放入到测试的函数中
           if !reflect.DeepEqual(res, v.want) {
             t.Errorf("期望找到: %v, 实际得到: %v\n", v.want, res)
           }
         })
       }
     }
     ```

   - 还可以使用 `go test` 中的 `-run` 选项来运行**指定函数中的指定子测试用例**。示例如下:

     ```go
     go test -run TestSplit/none
     go test -run 测试函数/子测试用例
     ```

4. > 测试覆盖率

   - 测试覆盖率是你的代码被测试套件覆盖的百分比。通常我们使用的都是语句的覆盖率，也就是在测试中至少被运行一次的代码占总代码的比例。

   - Go提供内置功能来检查你的代码覆盖率。我们可以使用 `go test -cover` 来查看测试覆盖率。例如:

     ```go
     $ go test -cover
     PASS
     coverage: 100.0% of statements
     ok    Go_Learn/Day_07/03_test_child/split     0.004s
     ```

     从上面的结果可以看到我们的测试用例覆盖了100%的代码。

   - Go还提供了一个额外的 `-coverprofile` 参数，用来将覆盖率相关的记录信息输出至一个文件。例如:

     ```go 
     $ go test -cover -coverprofile=test.out
     PASS
     coverage: 100.0% of statements
     ok    Go_Learn/Day_07/03_test_child/split     0.004s
     ```

     上面的命令会将覆盖率相关的信息输出到当前文件夹下面的 `test.out` 文件中，然后我们执行 `go tool cover -html=test.out` ，使用 `cover` 工具来处理生成的记录信息，该命令会打开本地的浏览器生成一个HTML报告。"绿色"标记的语句块表示被覆盖了，而"红色"的表示没有被覆盖。

5. > 深度判断

   - 使用reflect包中的 `DeepEqual` 方法，可以做到先判断数据类型是否一致，再判断里面的元素是否都一致，返回一个bool值

     ```go
     a := 123
     b := "123"
     if res := reflect.DeepEqual(a, b); res {
       fmt.Println("一样")
     } else {
       fmt.Println("不一样")
     }
     ```

6. > 基准测试 (Benchmark)

   - 基准测试就是在一定的工作负载之下检测程序性能的一种方法。基准测试的基本格式如下:

     ```go
     func BenchmarkName(b *testing.B) {
       // ...
     }
     ```

     基准测试以 `Benchmark` 为前缀，需要一个 `*testing.B` 类型的参数b，基准测试必须要执行 `b.N` 次，这样的测试才有对照性，`b.N` 的值是系统根据实际情况去调整的，从而保证测试的稳定性。`testing.B` 拥有的方法和 `testing.T` 基本一致，详情请看手册。

   - 根据 `子测试` 例子中的 `Split` 函数，基准测试代码如下

     ```go
     func BenchmarkSplit(b *testing.B) {
       for i := 0; i < b.N; i++ {
         Split("1,2,3", ",")
       }
     }
     ```

     基准测试并不会默认执行，需要增加 `-bench` 参数，通过执行 `go test -bench=Split` 命令执行基准测试，`-bench=` 后面的参数为正则表达式用于匹配函数名，输出结果如下:

     ```go
     $ go test -bench=Split
     goos: darwin
     goarch: amd64
     pkg: Go_Learn/Day_07/04_test_standard/split
     BenchmarkSplit-16        5171655               232 ns/op
     PASS
     ok      Go_Learn/Day_07/04_test_standard/split  1.439s
     ```

     其中 `BenchmarkSplit-16` 表示对`Split`函数进行基准测试，数字 `16` 表示 `GOMAXPROCS` 的值，这个对于并发基准测试很重要，可以通过后面跟 `-cpu=N` 来指定用几个CPU核心来测试，例如: `go test -bench=Split -cpu=1` 。`5171655` 和 `232 ns/op` 表示每次调用 `Split` 函数耗时 `232ns` ，这个结果是 `5171655` 次调用的平均值。

   - 我们还可以为基准测试添加 `-benchmem` 参数，来获得内存分配的统计数据

     ```go
     $ go test -bench=Split -benchmem
     goos: darwin
     goarch: amd64
     pkg: Go_Learn/Day_07/04_test_standard/split
     BenchmarkSplit-16        4998319       239 ns/op       112 B/op        3 allocs/op
     PASS
     ok      Go_Learn/Day_07/04_test_standard/split  1.444s
     ```

     其中，`112 B/op` 表示每次操作内存分配了112字节，`3 allocs/op` 则表示每次操作进行了3次内存分配。我们将 `Split` 函数优化一下如下:

     ```go
     func Split(str, sep string) (res []string) {
       res = make([]string, 0, strings.Count(str, sep)+1)
       index := strings.Index(str, sep)
       for index > -1 {
         res = append(res, str[:index])
         str = str[index+len(sep):]
         index = strings.Index(str, sep)
       }
       res = append(res, str)
       return
     }
     ```

     这一次提前使用make函数将res初始化为一个容量足够大的切片，而不再像之前一样通过调用append函数来追加。执行结果如下:

     ```go
     $ go test -bench=Split -benchmem
     goos: darwin
     goarch: amd64
     pkg: Go_Learn/Day_07/04_test_standard/split
     BenchmarkSplit-16       10760896               110 ns/op              48 B/op          1 allocs/op
     PASS
     ok      Go_Learn/Day_07/04_test_standard/split  1.307s
     ```

     这个使用make函数提前分配内存的改动，减少了2/3的内存分配次数，并且减少了一半的内存分配

7. > 性能比较函数

   - 上面的基准测试只能得到给定操作的绝对耗时，但是很多性能问题是发生在两个不同操作之间的相对耗时，比如同一个函数处理1000个元素的耗时与处理1万甚至100万个元素的耗时的差别是多少？再或者对于同一个任务使用哪种算法性能最佳？我们通常需要对两个不同算法的实现使用相同的输入来进行基准测试。

   - 性能比较函数通常是一个带有参数的函数，被多个不同的Benchmark函数传入不同的值来调用。举个例子如下:

     ```go
     func benchmark(b *testing.B, size int) { /* ... */ }
     func Benchmark10(b *testing.B) { benchmark(b, 10) }
     func Benchmark100(b *testing.B) { benchmark(b, 100) }
     func Benchmark1000(b *testing.B) { benchmark(b, 1000) }
     ```

     例如我们编写了一个计算 **斐波那契数列** 的函数如下:

     ```go
     // fib.go
     
     // Fib 是一个计算第n个斐波那契数的函数
     // 斐波那契数列: 后一个数是前两个数的和
     func Fib(n int) int {
       if n < 2 {
         return n
       }
       return Fib(n-1) + Fib(n-2)
     }
     ```

     我们编写的性能比较函数如下:

     ```go
     // fib_test.go
     
     func benchmarkFib(b *testing.B) {
       for i := 0; i < b.N; i++ {
         Fib(n)
       }
     }
     
     func BenchmarkFib1(b *testing.B) { benchmarkFib(b, 1) }
     func BenchmarkFib2(b *testing.B) { benchmarkFib(b, 2) }
     func BenchmarkFib3(b *testing.B) { benchmarkFib(b, 3) }
     func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
     func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
     func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }
     ```

     运行基准测试:

     ```go
     $ go test -bench=.
     goos: darwin
     goarch: amd64
     pkg: Go_Learn/Day_07/05_test_performance/fib
     BenchmarkFib1-16        508824984                2.25 ns/op
     BenchmarkFib2-16        216067983                5.56 ns/op
     BenchmarkFib3-16        126685104                9.37 ns/op
     BenchmarkFib10-16        3424062               347 ns/op
     BenchmarkFib20-16          27906             43381 ns/op
     BenchmarkFib40-16              2         651956174 ns/op
     PASS
     ok      Go_Learn/Day_07/05_test_performance/fib 10.457s
     ```

     这里需要注意的是，默认情况下，每个基准测试至少运行**1秒**。如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50, ...增加，并且函数再次运行。

     最终的BenchmarkFib40只运行了两次，每次运行的平均值只有不到一秒。像这种情况下我们应该使用 `-benchtime` 标志增加最小基准时间，以产生更准确的结果。例如:

     ```go
     goos: darwin
     goarch: amd64
     pkg: Go_Learn/Day_07/05_test_performance/fib
     BenchmarkFib1-16        1000000000               2.25 ns/op
     BenchmarkFib2-16        1000000000               5.49 ns/op
     BenchmarkFib3-16        1000000000               9.35 ns/op
     BenchmarkFib10-16       69106652               347 ns/op
     BenchmarkFib20-16         555812             43716 ns/op
     BenchmarkFib40-16             36         663695534 ns/op
     PASS
     ok      Go_Learn/Day_07/05_test_performance/fib 92.464s
     ```

     这一次 `BenchmarkFib40` 函数运行了36次，结果就会更准确一些。

     使用性能比较函数做测试的时候一个容易犯的错误就是把 `b.N` 作为输入的大小，例如以下两个例子都是错误的:

     ```go
     // 错误示范1
     func BenchmarkFibWrong1(b *testing.B) {
       for n := 0; b < b.N; n++ {
         Fib(n)
       }
     }
     
     // 错误示范2
     func BenchmarkFibWrong2(b *testing.B) {
       Fib(b.N)
     }
     ```

8. > 重置时间

   - `b.ResetTimer` 之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作。

     ```go
     func BenchmarkSplit(b *testing.B) {
       time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作, 比如链接数据库
       b.ResetTimer() 						  // 重置计时器
       for i := 0; i < b.N; i++ {
         Split("1.2.3", ".")
       }
     }
     ```

9. > 并行测试

   - `func (b *B) runParallel(body func(*PB))` 会以并行的方式执行给定的基准测试。

     `RunParallel` 会创建出多个 `goroutine`，并将 `b.N` 分配给这些 `goroutine` 执行，其中 `goroutine` 数量的默认值为 `GOMAXPROCS`。用户如果想要增加非CPU受限 (non-CPU-bound) 基准测试的并行性，那么可以在 `RunParallel` 之前调用 `SetParallelism`。`RunParallel` 通常会与 `-cpu` 标志一同使用。

     ```go
     func BenchmarkSplitParallel(b *testing.B) { // Parallel 用来这个函数为并行测试函数
       // b.SetParallelism(1) // 设置使用的CPU核心数
       b.RunParallel(func(pb *testing.PB) {
         for pb.Next() {
           Split("1,2,3", ",")
         }
       })
     }
     ```

     执行一下基准测试:

     ```go
     $ go test -bench=.
     goos: darwin
     goarch: amd64
     pkg: Go_Learn/Day_07/06_test_parallel/split
     BenchmarkSplitParallel-16       31082677                40.0 ns/op
     PASS
     ok      Go_Learn/Day_07/06_test_parallel/split  1.289s
     ```

     还可以通过在测试命令后添加 `-cpu` 参数如 `go test -bench=. -cpu 1 ` 来指定使用的CPU核心数量

10. > Setup与TearDown

   - 测试程序有时需要在**测试之前进行额外的设置 (setup) **或 **在测试之后进行拆卸 (teardown)**

   - TestMain

     如果测试文件包含函数: `func TestMain(m *testing.M)` 那么生成的测试会先调用 `TestMain(m)`，然后再运行具体测试。`TestMain` 运行在住 `goroutine` 中，可以在调用 `m.Run` 前后做任何 **设置 (setup)** 和 **拆卸(tardown)**。退出测试的时候应该使用 `m.Run` 的返回值作为参数调用 `os.Exit`。

     一个使用 `TestMain` 来设置Setup和TearDown的示例如下:

     ```go
     func TestMain(m *testing.M) {
       fmt.Println("write setup code here...") // 测试之前做的一些设置, 比如连接数据库
       // 如果 TestMain 使用了 flags, 这里应该加上 flag.Parse()
       retCode := m.Run()													// 执行测试
       fmt.Println("write teardown code here...")  // 测试之后做一些拆卸工作
       os.Exit(retCode)														// 退出测试
     }
     ```

     需要注意的是：在调用 `TestMain` 时，`flag.Parse` 并没有被调用。所以，如果 `TestMain` 依赖于command-line 标志 (包括testing包的标记) ，则应该显示的调用 `flag.Parse`。

11. > 子测试的Setup与Teardown

    - 有时候我们可能需要为每个测试集设置Setup与Teardown，也有可能需要为每个子测试设置Setup与Teardown。下面我们定义两个函数工具函数如下：

      ```go
      // 测试集的Setup与Teardown
      func setupTestCase(t *testing.T) func(t *testing.T) {
        t.Log("如果需要在此执行: 测试之前的Setup")
        return func(t *testing.T) {
          t.Log("如果需要在此执行: 测试之后的Teardown")
        }
      }
      
      // 子测试的Setup与Teardown
      func setupSubTest(t *testing.T) func(t *testing.T) {
        t.Log("如果需要在此执行: 子测试之前的Setup")
        return func(t *testing.T) {
          t.Log("如果需要在此执行: 子测试之后的teardown")
        }
      }
      ```

      使用方式如下:

      ```go
      func TestSplit(t *testing.T) {
        // 定义test结构体
        type test struct {
          str  string
          sep  string
          want []string
        }
      
        // 测试用例使用map存储
        tests := map[string]test{
          "one": {str: "1,2,3", sep: ",", want: []string{"1", "2", "3"}},
          "two": {str: "1,2,3", sep: ",", want: []string{"1,2,3"}},
        }
      
        teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
        defer teardownTestCase(t)            // 测试之后执行teardown操作
      
        for name, tc := range tests {
          t.Run(name, func(t *testing.T) { // 使用t.Run() 执行子测试
            teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
            defer teardownSubTest(t)           // 测试之后执行teardown操作
            got := Split(tc.str, tc.sep)
            if !reflect.DeepEqual(got, tc.want) {
              t.Errorf("want: %#v, got:%#v\n", tc.want, got)
            }
          })
        }
      }
      ```

      测试结果如下:

      ```go
      $ go test -v
      === RUN   TestSplit
      === RUN   TestSplit/one
      === RUN   TestSplit/two
      --- FAIL: TestSplit (0.00s)
          split_test.go:10: 如果需要在此执行: 测试之前的Setup
          --- PASS: TestSplit/one (0.00s)
              split_test.go:18: 如果需要在此执行: 子测试之前的Setup
              split_test.go:20: 如果需要在此执行: 子测试之后的Teardown
          --- FAIL: TestSplit/two (0.00s)
              split_test.go:18: 如果需要在此执行: 子测试之前的Setup
              split_test.go:48: want: []string{"1,2,3"}, got:[]string{"1", "2", "3"}
              split_test.go:20: 如果需要在此执行: 子测试之后的Teardown
          split_test.go:12: 如果需要在此执行: 测试之后的Teardown
      FAIL
      exit status 1
      FAIL    Go_Learn/Day_07/08_test_testmain_child/split    0.005s
      ```

12. > 示例函数

    - 被 `go test` 特殊对待的第三种函数就是示例函数，它们的函数名以 `Example` 为前缀。它们既没有参数也没有返回值。

      示例函数的格式

      ```go
      func ExampleName() {
        // ...
      }
      ```

      示例函数示例

      下面的代码使我们为 `Split` 函数编写的一个示例函数: 

      ```go
      func ExampleSplit() {
      	fmt.Println(Split("1,2,3", ","))
      	// Output:
      	// [1 2 3]
      }
      ```

      为代码编写示例函数有如下三个用处:

       - 示例函数能够作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联。

       - 示例函数只要包含了 `// Output:`  也是可以通过 `go test` 运行的可执行测试。

         ```go
         $ go test -run Example
         PASS
         ok      Go_Learn/Day_07/09_test_example/split   0.004s
         ```

       - 示例函数提供了可以直接运行的示例代码，可以直接在 `golang.org` 的 `godoc` 文档服务器上使用 `go Playground` 运行示例代码。

13. > socket编程

    - Socket是BSD UNIX的进程通信机制，通常也被称作"套接字"，用于描述IP地址和端口，是一个通信链的句柄。Socket可以理解为TCP/IP网络的API，它定义了许多函数或例程，程序员可以用它们来开发TCP/IP网络上的应用程序。电脑上运行的应用程序通常通过"套接字"向网络发出请求或这应答网络请求。

      **socket图解**

      `Socket` 是应用层与TCP/IP协议族通信的中间软件抽象层。在设计模式中，`Socket` 其实就是一个门面模式，它把复杂的TCP/IP协议族隐藏在`Socket`后面，对用户来说只需要调用Socket规定的相关函数，让 `Socket`去组织符合指定的协议数据然后进行通信。

      ![socket](https://www.liwenzhou.com/images/Go/socket/socket.png)
   
14. > Go语言实现TCP通信

    - **TCP协议**

      TCP/IP (Transmission Control Protocol/Internet Protocol) 即传输控制协议/网间协议，是一种面向连接 (连接导向) 的、可靠的、基于字节流的传输层 (Transport layer) 通信协议，因为是面向连接的协议，数据像水流一样传输，会存在黏包问题。

    - **TCP服务端**

      一个TCP服务器端可以同时连接很多个客户端，因为Go语言中创建多个goroutine实现并发非常方便和高效，所以我们可以每建立一次连接就创建一个goroutine去处理。

      TCP服务端程序的处理流程：

        1. 监听端口
        2. 接收客户端请求建立链接
        3. 创建goroutine处理链接

      - **TCP客户端**

        一个TCP客户端进行TCP通信的流程如下：

         1. 建立与服务端的链接 (拨号)
         2. 进行数据收发
         3. 关闭链接

15. > TCP粘包

    - 什么是粘包

      ```go
      收到client发来的数据： Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?
      收到client发来的数据： Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?
      收到client发来的数据： Hello, Hello. How are you?Hello, Hello. How are you?
      收到client发来的数据： Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?
      收到client发来的数据： Hello, Hello. How are you?Hello, Hello. How are you?
      ```

      比如说客户端分10次发送的数据，在服务端并没有成功的输出10次，而是多条数据"粘在了一起。

    - 为什么会出现粘包

      主要原因就是tcp数据传递模式是`流模式`，在保持长连接的时候可以进行多次的收和发。

      ”粘包“可发生在发送端也可发生在接收端：

       1. 由Nagle算法造成的发送端的粘包：Nagle算法是一种改善网络传输效率的算法。简单来说就是当我们提交一段数据给TCP发送时，TCP并不立刻发送此段数据，而是等待一小段时间看看在等待期间是否还有要发送的数据，若有则会一次把两段数据发送出去。
       2. 接收端接收不及时造成的接收端粘包，TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层取数据。当应用层由于某些原因不能及时的把TCP的数据取出来，就会造成TCP缓冲区中存放了几段数据。

     - 解决办法

       出现”粘包“的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行`封包`和`拆包`的操作。

       封包：封包就是给一段数据加上包头，这样一来数据包就分为包头和包体两部分内容了(过滤非法包时封包会加入”包尾“内容)。包头部分的长度是固定的，并且它存储了包体的长度，根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包。

       我们可以自己定义一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度。

16. > Go语言实现UDP通信

    - **UDP协议**

      UDP协议 (User Datagram Protocol) 中文名称是用户数据报协议，是OSI (Open System Interconnection, 开放式系统互联) 参考模型中一种无连接的传输层协议，不需要建立链接就能直接进行数据发送和接收，属于不可靠、没有时序的通信，但是UDP协议的实时性比较好。

    - **UDP服务端**

      UDP服务端程序的处理流程：

        1. 监听端口
        2. 接收客户端请求建立链接
        3. 创建goroutine处理链接

    - **UDP客户端**

      一个UDP客户端进行UDP通信的流程如下：

       1. 建立与服务端的链接
       2. 进行数据收发
       3. 关闭链接

