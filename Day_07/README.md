#### <center>Day_07</center>

1. 单元测试
2. 单元测试组
3. 子测试

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

   - 示例

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

   - 子测试还可以使用 `go test` 中的 `-run` 选项来运行**指定函数中的指定测试用例**。示例如下:
   - 

4. > 深度判断

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

5. 

