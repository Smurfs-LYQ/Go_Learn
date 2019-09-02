#### <center>Day06</center>

1. 接口
    * one 接口定义
    * two 接口实例
    * thr 接口嵌套
	* four 类型断言

#### <center>笔记</center>
1. > 接口
	1. interface类型默认是一个指针
	2. interface类型可以定义一组方法，但是这些不需要实现。并且interface不能包含任何变量
	3. Golang中的接口，不需要显性的实现。只要一个变量，含有接口类型中的“所有方法”，那么这个变量就实现了这个接口。因此，Golang中没有implement类似的关键字
	4. 如果一个变量含有了多个interface类型(接口)的方法，那么这个变量就实现了多个接口
	5. 如果一个变量只含有一个interface的部分方法, 那么这个变量就没有实现这个接口
2. > 定义:
	type 接口名 interface {
		方法名1(参数列表) 返回值列表
		方法名2(参数列表) 返回值列表
	}
3. > 多态
	* 字面意思即“多种形态”。一个事物的多种形态，都可以按照统一的接口进行操作即为多态。在面向对象语言中，接口的多种实现方式即为多态。
	车 可以代表汽车和火车等，所以 车 这个概念就是典型的多态，因为他不仅可以代表汽车 也可以代表火车等
	* 例如我们下面定义了一个T1的接口，其中有两个方法。后面又分别定义了Man和WoMan的结构体，并且他们都实现了T1接口中定义的两个方法。那么接口T1既可以通过Man结构体来实现自己，也可以通过WoMan结构体来实现自己，这种有多种实现方式的接口即为多态
4. > 类型断言
	* 由于接口是一般类型，不知道具体类型，如果要转成具体类型可以采用以下方法进行转换：
		```
		var t int 				   var t int
		var x interface{} 		   var x interface{}
		x = t 					   x = t
		y = x.(int)	// 转成int	    y, ok = x.(int) // 转成int, 带检查
		```