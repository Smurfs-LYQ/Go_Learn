#### <center>Day_03</center>

1. 指针
2. new 和 make
3. panic 和 recover
4. 自定义类型和类型别名
5. struct 结构体
6. 结构体的实例化
7. 构造函数和析构函数
8. init

#### <center>笔记</center>
1. > 指针
    - `指针` 和 `地址` 有什么区别?
    	- `地址` : 就是内存地址(用字节编码来描述的内存地址)
    	- `指针` : 指针是保存地址的变量
  	- `&` 和 `*`
    	- `&` : 表示取地址
    	- `*` : 表示根据地址取值
2. > new 和 make
	- `new` 是用来初始化值类型指针的
		```
		// 定义一个int类型的指针
		var T1 *int
		// 初始化int类型的指针(因为int是值类型所以用new)
		T1 = new(int)
		// 给T1赋值
		*T1 = 1
		```
	- `make` 是用来初始化引用类型的
3. > panic 和 recover
	- `panic` 是代码运行时的错误
	- `recover` 可以捕获 `panic` 报错，并且尝试将函数从当前的异常状态恢复
4. > 自定义类型和类型别名
	- 自定义类型
		```
		// NewInt 自定义类型: 基于Go语言的int类型，创建一个NewInt类型
		type NewInt int
		```
	- 类型别名
		```
		// MyInt 类型别名: 只存在代码编写过程中，代码编译之后根本不存在MyInt类型
		type MyInt = int
		```
5. > struct 结构体
	- 我们可以通过 `struct` 来定义自己的类型
	- 结构体的定义
		```
		// 使用 `type` 和 `struct` 关键字来定义结构体

		type 结构体类型名 struct {
			字段名 字段类型
			字段名 字段类型
			...
		}
		```
6. > 结构体的实例化
	- 基本实例化
		```
		// 如果初始化时没有给字段设置值，则默认使用对应类型的零值
		var 变量名 结构体类型名
		```
	- 实例化结构体，但其为结构体类型指针
		```
		var 变量名 new(结构体类型名)
		var 变量名 &结构体类型名
		```
  	- 结构体初始化
		```
		// 不省略字段名
		var 变量名 = 结构体类型名 {
			字段名: 值,
			字段名: 值,
			...
		}

		// 省略字段名，注意省略字段名的话需要按照结构体中字段声明顺序进行赋值
		var 变量名 = 结构体类型名 {
			值_1,
			值_2,
			...
		}
		```
7. > 构造函数和析构函数
	- Go语言的结构体没有 `构造函数` 和 `析构函数` ，但是我们可以根据Go语言的一些函数自己实现 `构造函数` 和 `析构函数`。
	- `构造函数`
    	- 析构函数的作用: 主要用来在创建对象时初始化对象
			```
			package main

			import "fmt"

			// 定义一个结构体
			type T1 struct {
				Name string
				Age  int
			}

			// 定义一个初始化结构体的方法
			/*
			创建T1类型的变量时，直接调用Test_1函数
			*/
			func Test_1(name string, age int) *T1 {
				return &T1{
					name,
					age,
				}
			}

			func main() {
				// 调用自定义析构函数
				one := Test_1("Smurfs", 21)
				fmt.Println(one)
			}
			```
	- `析构函数`
    	- 析构函数的作用: 在程序结束时自动执行(销毁一个对象)
    	- 可以使用 `defer` 来实现析构函数
8. > init
	- 在程序开始时自动执行
		```
		package main

		import "fmt"

		func init() {
			// 程序开始时自动执行
			fmt.Println("这个是init函数")
		}

		func main() {
			fmt.Println("这个是main函数")
		}
		```
9.  