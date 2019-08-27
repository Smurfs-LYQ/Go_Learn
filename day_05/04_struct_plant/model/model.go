package model

// 创建一个结构体
type man struct {
	Name string
	Age  int
}

// 创建一个函数, 接收两个函数, 返回一个Man结构体
func NewMan(name string, age int) *man {
	return &man{
		Name: name,
		Age:  age,
	}
}
