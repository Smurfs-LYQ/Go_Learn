package four

import "fmt"

/*
类型断言
*/

type Student struct {
	Name string
	Sex  string
}

// 空接口可以接受任何类型的数据
func T1(a interface{}) {
	b, ok := a.(int) // 将 b 转换为int类型，OK为转换的结果，返回bool值
	if ok == false {
		fmt.Println("没有转换成功")
		return
	}
	b += 3
	fmt.Println(b)
}

// 因为空接口可以接受任何类型的数据，而且这里还加了...所以可以接受任意数量任意类型的参数，并保存到对应的变量slice中
func T2(items ...interface{}) {
	for k, v := range items {
		switch v.(type) { // 这里面的type是关键词，这里是用来判断v的类型的
		case bool:
			fmt.Printf("这个是bool类型 %v %T\n", k, v)
		case float32, float64:
			fmt.Printf("这个是float类型 %v %T\n", k, v)
		case int:
			fmt.Printf("这个是int类型 %v %T\n", k, v)
		case string:
			fmt.Printf("这个是string类型 %v %T\n", k, v)
		case Student:
			fmt.Printf("这个是Student类型 %v %T\n", k, v)
		default:
			fmt.Printf("没列出来，检测一下: %T\n", v)
		}
	}
}

func Four() {
	var one string
	T1(one)

	fmt.Println()

	var two Student = Student{
		Name: "Smurfs",
		Sex:  "男",
	}
	T1(two)

	fmt.Println()

	T2(&two, "123", 123, 123.0)
}
