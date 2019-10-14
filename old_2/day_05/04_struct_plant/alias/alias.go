package alias

import "fmt"

// 给int数据类型设置别名为integer
type integer int

// 结构体是用户单独定义的类型, 不能和其他类型进行强制转换
type T1 struct {
	Number int
}

type Test_1 T1 // 给结构体T1设置一个别名, 但和T1并不是同一类型

func Alias_1() {
	var a T1
	a = T1{30}

	var b Test_1
	b = Test_1{20}

	// 结构体和其别名 都是用户单独定义的类型, 不能和其他类型进行强制转换
	// a = b
	a = T1(b)
	fmt.Println(a)

	var i integer = 1000
	var j int = 10

	// 虽然设置了别名, 但是依然不能直接进行int和integer的 = 操作
	// j = i
	j = int(i) // 需要将类型强制转换

	fmt.Println(i, j)
}
