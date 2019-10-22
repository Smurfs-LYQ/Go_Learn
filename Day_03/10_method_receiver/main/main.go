package main

import "fmt"

type T1 struct {
	Name string
	Age  int
	Sex  string
}

// 接受者是值类型，所以每次修改的只是对象的值拷贝
func (t T1) Test_1(name string, age int, sex string) *T1 {
	t.Name = name
	t.Age = age
	t.Sex = sex
	return &t
}

// 接受者是引用类型，所以每次修改的是对象本身
func (t *T1) Test_2(name string, age int, sex string) *T1 {
	t.Name = name
	t.Age = age
	t.Sex = sex
	return t
}

func main() {
	// 初始化一个结构体
	var T1_obj = T1{
		"Smurfs",
		21,
		"男",
	}
	fmt.Printf("%p, %v\n", &T1_obj, T1_obj)

	// 因为Test_1方法的接受者并不是T1_obj的指针，所以修改的只是T1_obj的值拷贝
	T1_obj.Test_1("Smurfs_1", 21, "男")
	fmt.Printf("%p, %v\n", &T1_obj, T1_obj)

	// 因为Test_2方法的接受者是T1_obj的指针，所以修改的是T1_obj其本身
	// (&T1_obj).Test_2("Smurfs_3", 21, "男")
	// 语法糖(简写)
	T1_obj.Test_2("Smurfs_3", 21, "男")
	fmt.Printf("%p, %v\n", &T1_obj, T1_obj)
}
