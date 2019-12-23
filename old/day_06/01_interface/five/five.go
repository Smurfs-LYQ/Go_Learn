package five

import "fmt"

// 创建一个接口
type T1 interface {
	// 添加方法
	Get_Name() string
}

// 创建一个结构体
type Struct_One struct {
	// 定义结构体属性
	Name string
}

// 创建一个属于Struct_One结构体的函数
func (this Struct_One) Get_Name() string {
	return this.Name
}

func Five() {
	// 实例化一个结构体
	var one Struct_One = Struct_One{Name: "Smurfs"}
	// 实例化一个接口
	var two T1
	// 给接口绑定一个实现了他所有方法的结构体
	two = one

	res, ok := two.(T1)
	fmt.Println(res, ok)

	// 这里讲的不详细
	// if sv, ok := one.(T1); ok { // 如果ok返回值不是nil则执行if里面的代码
	// 	fmt.Println(ok)
	// 	fmt.Printf("v implements T1():%s\n", sv.T1())
	// }
}
