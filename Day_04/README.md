#### <center>Day_04</center>

1. defer经典案例
2. Go语言中的包
3. init初始化函数
4. time包
5. Go语言基础之接口
6. 接口的嵌套
7. 空接口
8. 类型断言
9. 值接收者和指针接收者实现接口的区别
10. 文件操作-读取文件
11. 文件操作-bufio
12. 文件操作-ioutil
13. 文件操作-写入文件

#### <center>笔记</center>
1. > return原理和defer执行时机
	- 函数中return语句底层实现原理
		- `return x` = `返回值=x` => `RET指令`
	- defer语句执行的时机
		- `return x` = `返回值=x` => `运行defer` => `RET指令`
2. > Go语言中的包
	- 在工程化的Go语言开发项目中，Go语言的源码复用是建立在(package)基础之上的。
	- 包介绍
    	- `包(package)` 是多个Go源码的集合，是一种高级的代码复用方案，Go语言为我们提供了很多内置包，如 `fmt` 、 `os` 、 `io` 等。
  	- 定义包
    	- 一个包可以简单理解为一个存放 `.go` 文件的文件夹。该文件夹下面的所有go文件都要在代码的第一行添加如下代码，声明该文件归属的包。
			```
			package 包名
			```
    	- **注意事项**
        	- 一个文件夹下面只能有一个包，同样一个包的文件不能再多个文件夹下
        	- 包名可以不和文件夹的名字一样，包括不能包含`"-"符号`
        	- 包名为main的包为应用程序的入口包，编译时不包含main包的源代码时不会得到可执行文件
  	- 可见性
    	- 如果想在一个包中引用另一个包里的标签(如常量、变量、类型、函数等)，该标识符必须是对外可见的(public)。在Go语言中只需要将标识符的首字母大学就可以让标识符对外可见了。
  	- 包的导入
		```
		// 单个包导入
		import "包的路径"

		// 多个包导入
		import (
			"包的路径"
			"包的路径"
			...
		)
		```
		- **注意事项**
    		- 包的路径格式为 **`GOPATH/src/`** 后面的路径
  	- 自定义包名
		```
		// 单个包导入
		import 自定义包名 "包名"

		// 多个包导入
		import (
			自定义包名 "包名"
			自定义包名 "包名"
			...
		)
		```
  	- 匿名导入包
    	- 如果只希望导入包，而不使用包内部的数据时，可以使用匿名导入包。
			```
			import _ "包的路径"
			```
    	- 匿名导入的包与其他方式导入的包一样都会被编译到可执行文件中。
3. > init初始化函数
	- Go语言程序执行导入包语句会自动触发包内部 `init()` 函数的调用。
	- **注意事项**
    	- `init()` 函数没有参数也没有返回值。
    	- `init()` 函数在程序运行时自动被调用执行，不能在代码中主动调用它。
    	- 如果导入包只想执行其 `init()` 函数，那可以使用匿名导入包。
  	- 包中 `init` 函数的执行时机
    	- 全局声明(例如定义变量或常量)
    	- init()
    	- main()
4. > time包
	- 时间类型
	- `time.Time` 类型表示时间
		```
		func timeDemo() {
			now := time.Now() // 获取当前时间
			fmt.Printf("%T\n%v\n", now, now)

			fmt.Println(now.Year())   // 年
			fmt.Println(now.Month())  // 月
			fmt.Println(now.Day())    // 日
			fmt.Println(now.Hour())   // 小时
			fmt.Println(now.Minute()) // 分钟
			fmt.Println(now.Second()) // 秒
		}
		```
	- 时间戳
		```
		// 时间戳是自1970年1月1日(08:00:00GMT)至当前时间的总毫秒数。它也被称为Unix时间戳(UnixTimestamp)。

		func timestampDemo1() {
			now := time.Now()           // 获取当前时间
			fmt.Println(now.Unix())     // 时间戳
			fmt.Println(now.UnixNano()) // 纳秒时间戳
		}
		```
	- 使用 `time.Unix()` 函数将时间戳转为时间格式。
		```
		func timestampDemo2() {
			now := time.Now()                   // 获取当前时间
			UnixTimestamp := now.Unix()         // 获取当前时间戳
			time := time.Unix(UnixTimestamp, 0) // 将当前时间戳转换为时间格式
			fmt.Println(time)
			fmt.Println(time.Year())   // 年
			fmt.Println(time.Month())  // 月
			fmt.Println(time.Day())    // 日
			fmt.Println(time.Hour())   // 时
			fmt.Println(time.Minute()) // 分
			fmt.Println(time.Second()) // 秒
			
			fmt.Printf("%4d-%2d-%0d %02d:%02d:%02d\n", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second())
		}
		```
	- 定时器
	- 使用 `time.Tick(时间间隔)` 来设置定时器。
		```
		func tickDemo() {
			ticker := time.Tick(time.Second) // 定义一个1秒间隔的定时器
			for i := range ticker {
				fmt.Println(i) // 每秒都会执行的任务
			}
		}
		```
		- 时间间隔
			```
			// Duration类型代表两个时间点之间经过的时间，以纳秒为单位。可表示的最长时间段大约290年。定义的时间间隔常量如下:
			const (
				Nanosecond  Duration = 1					// 纳秒
				Microsecond			 = 1000 * Nanosecond	// 微妙
				Millisecond			 = 1000 * Microsecond	// 毫秒
				Second				 = 1000 * Millisecond	// 秒
				Minute				 = 60 * Second			// 分钟
				Hour				 = 60 * Minute			// 小时
			)
			// 例如: `time.Duration` 表示1纳秒，`time.Second` 表示1秒。
			```
	- 时间格式化
    	- 时间类型有一个自带的方法 `Format` 进行格式化，需要注意的是Go语言中格式化时间模板不是常见的 `Y-m-d H:M:S` 而是使用Go的诞生时间2006年1月2号15点04分(记忆口诀为20061234)。
			```
			func formatDemo() {
				fmt.Println("################时间格式化################")
				now := time.Now()
				// 格式化的模板为Go的出生时间2006年1月2号15点04分
				fmt.Println(now.Format("2006-01-02 15:04"))
				fmt.Println(now.Format("2006/01/02 15:04"))
				fmt.Println(now.Format("15:04 2006/01/02"))
				fmt.Println(now.Format("2006/01/02"))
				
				fmt.Println(now.Format("2006-01-02 15:04:05.000")) // 05表示秒 .000表示毫秒
			}
			```
		- 如果想格式化为12小时方式，需指定 `PM`，并且小时由`15`改为`03`
			```
			func formatDemo() {
				// 格式化为12小时制格式
				fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
				// PM  代表12小时制格式
				// Mon 代表周几
				// Jan 代表几月
			}
			```
  	- 时间加时间间隔
    	- 我们在日常的编码过程中可能会遇到要求 时间+时间间隔 的要求，Go语言的时间对象提供了一个Add方法
  			```
			func (t Time) Add(d Duration) Time
  			```
    	- 参数:
        	- Duration time包的时间间隔常量
  	- 两个时间相减
    	- 求两个时间之间的差值,Go语言的时间对象提供了一个Sub方法
  			```
			func (t Time) Sub(u Time) Duration
			```
        	- 参数
            	- 两个时间相减的另一个时间对象
        	- 返回值
            	- 返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，将返回最大值/最小值。要获取时间点t-d(d为Duration)，可以使用t.Add(-d)
  	- 时间比较
    	- Equal
			```
			func (t Time) Equal(u Time) bool
			```
        	- 判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息
    	- Before
			```
			func (t Time) Before(u Time) bool
			```
        	- 如果t代表的时间点在u之前，则返回真；否则返回假
    	- After
			```
			func (t Time) After(u Time) bool
			```
        	- 如果t代表的时间点在u之后，则返回真；否则返回假
5. > Go语言基础之接口
	- 接口`(interface)`定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。
	- 接口介绍
		- 在Go语言中的接口(interface)是一种**类型**，一种抽象的类型。
		- `interface`是一组`method`的集合，是`duck-type programming`的一种体现。接口做的事情就像是定义一个协议(规则)，只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。不关心属性(数据)，只关心行为(方法)。
  	- 接口的定义
    	- Go语言提倡面向接口编程。
    	- 每个接口由数个方法组成，接口的定义格式如下：
			```
			type 接口类型名 interface {
				方法1( 参数列表1 ) 返回值列表1
				方法2( 参数列表2 ) 返回值列表2
				...
			}
			```
    	- 其中
        	- 接口名: 使用`type`将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加`er`，如有写操作的接口叫`Writer`，有字符串功能的接口叫`Stringer`等。
        	- 方法名: 当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包`(package)`之外的代码访问。
        	- 参数列表、返回值列表: 参数列表和返回值列表中的参数变量名可以省略。
  	- 接口的实现
    	- 只要对象实现了接口中的规则就实现了接口
  	- 一个类型实现多个接口
    	- 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。例如，狗可以归为动物类别，在细分下去可以归为哺乳动物类别。
    	- 并且一个接口的方法，不一定需要由一个类型完全实现，接口的方法也可以通过在类型中切入其他类型或结构体来实现。
  	- 接口类型变量
    	- 接口类型变量能够存储所有实现了该接口的实例。
  	- 接口值
    	- 接口值为引用类型
    	- 接口值由两部分组成
        	- 动态数据类型
        	- 对应值得内存地址
6. > 接口嵌套
	- 接口与接口之间可以通过嵌套创造出新的接口，操作和结构体的嵌套类似。
7. > 空接口
	- 空接口的定义
		- 空接口是指没有定义任何方法的接口，因此任何类型都实现了空接口。
    - 空接口类型的变量可以存储任意类型的变量。
		```
		// 声明一个空接口
		var T1 interface{}
		// 因为任何类型都实现了空接口，所以任何类型都可以对空接口进行赋值
		T1 = 100
		fmt.Printf("%T %v\n", T1, T1)
		T1 = "Smurfs"
		fmt.Printf("%T %v\n", T1, T1)
		T1 = false
		fmt.Printf("%T %v\n", T1, T1)
		```
  	- 空接口的应用
    	- 空接口作为函数的参数
			```
			// 使用空接口实现可以接受任意类型的函数参数
			// T1 声明一个可以接收任何类型参数的函数
			func T1(a interface{}) {
				fmt.Printf("%T %v\n", a, a)
			}
			```
    	- 空接口作为map的值
			```
			// 使用空接口实现可以保存任意值的字典
			var map_1 = make(map[int]interface{})
			map_1[0] = 123
			map_1[1] = struct{ name string }{name: "Smurfs"} // 传一个结构体也可以接收
			map_1[2] = false
			for k, v := range map_1 {
				fmt.Printf("%d %T %v\n", k, v, v)
			}
			```
8. > 类型断言
	- 空接口可以存储任意类型的值，类型断言可以帮我们判断接收到的值是不是我们想要`类型`的值。
	- 用于if判断的
		```
		语法格式: 
			x.(T)
		其中:
			1. x : 表示类型为"interface{}"的变量
			2. T : 表示断言"x"可能是类型
		```
		- 该语法返回两个值
    		- 第一个值 如果第二个值为`true`则为`x`的值类型为`T`, 如果第二个值为`false`则为`T`类型对应的零值
    		- 第二个值是一个布尔值，若为`true`则表示断言成功，为`false`则表示断言失败
   	- 用于switch判断的
		```
		语法格式:
			x.(type)
		其中:
			1. x : 表示类型为"interface{}"的变量
			2. type : 固定格式，他可以表示所有类型
		示例:
			switch x.(type) {
			case int :
				fmt.Println("is int")
			case string :
				fmt.Println("is string")
			...
			}
		```
		该语法也返回两个参数  
		第一个值 如果第二个值为`true`则为`x`的值类型为`T`, 如果第二个值为`false`则为`T`类型对应的零值  
		第二个值 是一个布尔值，若为`true`则表示断言成功，为`false`则表示断言失败  
		但是第二个参数被switch接收了，所以不用额外命名
9. > 值接收者和指针接收者实现接口的区别
	- 使用值接收者实现的方法，无论是用`对象的值还是指针`都可以赋值给`接口变量`，因为Go语言中有对指针类型变量求值的语法糖，而且指针指向的是对象本身，所以并不影响。
	- 使用指针接收者实现的方法，只能用`对象的指针`赋值给`接口变量`，因为`对象的值`无法直接获取到它的内存地址。
10. > 文件操作-读取文件
	- 打开文件
    	- `os.open()` 函数能够打开一个文件，返回一个 `*file` 和 `err`。
	- 关闭文件
    	- 对得到的文件实例调用 `close()` 方法能够关闭
    	- 通常为了防止忘记关闭文件都会配合 `defer` 来使用
  	- 打开和关闭文件
		```
		// 打开文件
		file, err := os.Open("./test.txt")
		// 关闭文件
		defer file.Close()
		if err != nil {
			fmt.Println("打开文件失败, 错误信息: ", err)
			return
		}
		fmt.Println(file)
		```
  	- 读取文件
    	- file.Read()
        	- 它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回 `0` 和 `io.EOF`。
11. > 文件操作-bufio
	- `bufio`是在file的基础上封装了一层API，支持更多的功能
12. > 文件操作-ioutil
	- `io/ioutil` 包的 `ReadFile` 方法能够读取完整的文件，只需要将文件名作为参数传入，返回值(读取到内容的byte切片、错误信息)
13. > 文件操作-写入文件
	- `os.OpenFile()` 函数能够以指定模式打开文件，从而实现文件写入相关功能
    	- 参数 1: 要打开的文件名
    	- 参数 2: 打开文件的模式
    	- 参数 3: 文件权限，一个八进制数。
        	- r (读) 04
        	- w (写) 02
        	- x (执行) 01
	- 打开文件的模式:
		| 运算符 | 含义 |
		| ----- | --- |
		| os.O_WRONLY   | 只写 |
		| os.O_CREATE   | 创建文件 |
		| os.O_RDONLY   | 只读 |
		| os.O_RDWR     | 读写 |
		| os.O_TRUNC    | 清空 |
		| os.O_APPEND   | 追加 |
  	- `Write` 和 `WriteString`
  	- `bufio.NewWriter`
  	- `ioutil.WriteFile`
14. > os.Stat()
	- 作用: 获取文件的信息
	- 参数
    	- 文件的路径
  	- 返回值
    	- 文件信息的结构体 FileInfo
    	- 错误信息
15. > fmt.Fprintf()
	- 这是一个写入文件用的格式化参数
	- 参数:
    	-  io.Writer 		一个可以写入的文件类型对象
    	-  string	 		Printf()函数中格式化的语句，例: "用户ID: %d"	可以省略
    	-  ...interface{}	与第二个参数对应的变量						    可以省略
	- 返回值:
    	- 写入的字节数
    	- 错误信息
16. > runtime 包
	- 他可以获取到程序在执行的时候环境的互操作，如正在执行的Go函数。
17. > path 包
	- 这个包是专门用来操作路径信息的。
18. 