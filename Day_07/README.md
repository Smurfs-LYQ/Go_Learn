#### <center>Day_07</center>

1. 单元测试
2. 单元测试组
3. 子测试
4. 基准测试

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

6. > 基准测试 (BenchmarkXxx)

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

     基准测试并不会默认执行，需要增加 `-bench` 参数，通过执行 `go test -bench=Split` 命令执行基准测试，`-bench=` 后面的参数就是函数名 `Benchmark` 后面跟的内容，输出结果如下:

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

8. > 

   - 
