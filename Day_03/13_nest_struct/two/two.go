package two

import "fmt"

// 结构体嵌套匿名结构体

type T1 struct {
	Name string
	Age  int
}

type T2 struct {
	T1  // 嵌套匿名结构体
	Age int
	Sex string
}

func Two() {
	var Test_1 = T2{
		T1{
			"Smurfs",
			21,
		},
		20,
		"男",
	}
	fmt.Println(Test_1)

	// 修改Test_1中嵌套结构体中的值
	// 修改Test_1和嵌套结构体不重复的字段
	// Test_1.T1.Name = "Smurfs的格格巫"
	Test_1.Name = "Smurfs的格格巫"
	// 修改Test_1和嵌套结构体中重复的字段
	Test_1.Age = 21    // 默认修改的是Test_1自身的字段
	Test_1.T1.Age = 20 // 修改嵌套结构体中重复的字段需要加上嵌套结构体的类型名或者变量名
	fmt.Println(Test_1)
}
