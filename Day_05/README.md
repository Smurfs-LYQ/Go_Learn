#### <center>Day_05</center>

1. 反射-TypeOf
2. 反射-ValueOf
3. 反射-Elem


#### <center>笔记</center>
1. > 反射
	- 反射就是在运行时动态的获取一个变量的类型信息和值信息。
	```
  	反射是指程序运行期间对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。
	支持反射的语言可以再程序编译期间将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期间获取类型的反射信息，并且有能力修改它们。
	Go程序在运行期间使用reflect包访问程序的反射信息。
	```
	- 反射的作用和类型断言比较相似，都是用来判断变量的类型的。但是Go语言的类型太多，类型断言猜不全，使用反射就能直接拿到接口值的动态类型和动态值。
	- 优点
    	- 让代码更灵活
	- 缺点
    	- 执行效率低	 
2. > reflect包
 	- 在Go语言的反射机制中，**任何接口值都是由** `一个具体类型` 和 `具体类型的值` 两部分组成的。在Go语言中反射的相关功能由内置的reflect包提供，任何接口值在反射中都可以理解为由 `reflect.Type` 和 `reflect.Value` 两部分组成，并且reflect包提供了 `reflect.TypeOf` 和 `reflect.ValueOf` 两个函数来获取人以对象的Value和Type。
 	- **TypeOf**
     	- 在Go语言中，使用 `reflect.TypeOf` 函数可以获取任意值得类型对象(reflect.Type)，程序通过类型对象可以访问任意值的类型信息。
			```
			func reflectType(x interface{}) {
				v := reflect.TypeOf(x)		// 可以拿到x的动态类型信息
				fmt.Printf("%T %v\n", x, x) // 原理就是用的反射，代码补全的原理也是反射
				fmt.Printf("%T %v\n", v, v)
			}
			```
   	- **type name** 和 **type kind**
     	- 在反射中关于类型还划分为两种: `类型(type)` 和 `种类(kind)`。 因为在Go语言中我们可以使用type关键字构造很多自定义类型，而 `种类(kind)` 就是指底层的类型，但在反射中，当需要划分指针、结构体等大品种的类型时，就会用到 `种类(kind)`。 举个例子，我们定义了两个指针类型和两个结构体类型，通过反射查看它们的类型和种类。
			```
			func reflectType(x interface{}) {
				t := reflect.TypeOf(x)      // 可以拿到x的动态类型信息
				fmt.Printf("type:%v name:%v kind:%v\n", t, t.Name(), t.Kind()) // 原理就是用的反射，代码补全的原理也是反射
			}

			// Cat 定义一个猫的结构体
			type Cat struct {
				name string
			}

			// People 定义一个人的结构体
			type People struct {
				name string
				age  int
			}

			func main() {
				var a *float32
				reflectType(a)
				Tom := Cat{"tom"}
				reflectType(Tom)
				Man := People{"Smurfs的格格巫", 18}
				reflectType(Man)
			}
			```
		- Go语言的反射中对所有指针变量的 `种类(kind)` 都是 `ptr`，但需要注意的是指针变量而类型 `名称(Name)` 是 `空`。
 	- **ValueOf**
     	- `reflect.ValueOf()` 返回的是 `reflect.Value` 类型，其中包含了原始值的值信息。`reflect.Value` 与原始值之间可以互相转换。
     	- `reflect.Value` 类型提供的获取原始值的方法如下:

		|   方法   			|  说明  |
		|   ---   			|  ---  |
		| Interface() interface{} | 将值以interface{}类型返回，可以通过类型断言转换为指定类型 |
		| Int() int64 		| 将值以int类型返回，所有有符号整型均可以此方式返回 |
		| Uint() uint64 	| 将值以uint类型返回，所有无符号整型均可以此方式返回 |
		| Float() float64 	| 将值以双精度(float64)类型返回，所有浮点数(float32、float64)均可以此方式返回 |
		| Bool() bool 		| 将值以bool类型返回 |
		| Bytes() []bytes 	| 将值以字节数组[]bytes类型返回 |
		| String() string 	| 将值以字符串类型返回 |
		- 通过反射获取值
			```
			func reflectValue(x interface{}) {
				v := reflect.ValueOf(x) // 获取接口的值信息
				k := v.Kind()           // 拿到值对应的种类
				fmt.Printf("%#v, %v\n", v, k)
				switch k {
				case reflect.Int:
					// v.Int() 从反射中获取整型的原始值，然后通过int()强制类型转换
					fmt.Printf("这是一个int类型的值，值为: %d\n", int(v.Int()))
				case reflect.String:
					// v.String() 从反射中获取字符串的原始值，然后通过string()强制类型转换
					fmt.Printf("这是一个string类型的值，值为: %s\n", string(v.String()))
				case reflect.Bool:
					// v.Bool() 从反射中获取布尔值的原始值，然后通过bool()强制类型转换
					fmt.Printf("这是一个bool类型的值，值为: %t\n", bool(v.Bool()))
				default:
					fmt.Println(k)
				}
			}

			func main() {
				reflectValue(1234)
				reflectValue("1234")
				reflectValue(false)
				reflectValue(12.34)

				// 将int类型的原始值转换为revlect.Value类型
				fmt.Printf("type : %T\n", reflect.ValueOf(10))
			}
			```
		- 通过反射设置变量的值
    		- 想要在函数中通过反射修改变量的值，需要注意函数参数传递的是**值拷贝**，必须传递变量地址才能修改变量值。而反射中使用专有的 `Elem()` 方法来获取指针对应的值
				```
				// 通过反射修改值
				func reflectElem(x interface{}) {
					// 首先将x转换为reflect.Value类型的变量v
					v := reflect.ValueOf(x)

					// 通过reflect包中的Kind方法获取变量v的类型
					if v.Kind() == reflect.Ptr {
						/*
							修改变量v的值
								因为函数传参都是值传递，所以如果想修改其本身的话需要将其对应的内存地址传入
								但是函数接收的是一个interface类型，所以通过将其转换为reflect.Value类型
								然后再调用reflect的Elem函数获取他内存指向的值，之后再调用reflect中设置
								值得函数就可以了(例如: SetInt SetString...)
						*/
						v.Elem().SetInt(123)
					}
				}

				func main() {
					var a int64 = 100
					reflectElem(&a)
					fmt.Println(a)
				}
				```
  		- isNil()
			```
			func (v Value) IsNil() bool
			```
    		- `IsNil()` 导致v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。
  		- isValid()
			```
			func (v Value) IsValid() bool
			```
			- `IsValid()` 返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。