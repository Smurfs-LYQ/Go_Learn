package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name" smurfs:"名字"` // 多个Tag使用"空格"分隔
	Score int    `json:"score" smurfs:"成绩"`
}

func (s student) SetStudent(name string, score int) {
	s.Name = name
	s.Score = score
	fmt.Println("设置成功", s)
}

func (s student) GetStudent() {
	fmt.Printf("姓名: %s\n", s.Name)
	fmt.Printf("分数: %d\n", s.Score)
}

func reflectStructFunc(name string, score int) {

}

func main() {
	stu1 := student{
		Name:  "Smurfs",
		Score: 99,
	}

	t := reflect.TypeOf(stu1)       // 获取变量stu1的反射对象类型信息
	v := reflect.ValueOf(stu1)      // 获取变量stu1的反射对象值信息
	fmt.Println(t.Name(), t.Kind()) // 获取结构体的类型以及底层类型

	// 通过for循环遍历结构体的所有字段信息
	fmt.Printf("\n############ 通过for循环遍历结构体的所有字段信息 ############\n")
	for i := 0; i < t.NumField(); i++ { // NumField() 通过反射获取结构体中所有的字段信息
		field := t.Field(i) // Field() 返回索引对应的结构体字段的信息
		fmt.Printf("name:%s index:%d type:%v tag:%v\n", field.Name, field.Index, field.Type, field.Tag)
	}

	// 通过字段名获取指定结构体字段信息
	fmt.Printf("\n############ 通过for循环遍历结构体的所有字段信息 ############\n")
	if scortField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v Smurfs_tag:%v\n", scortField.Name, scortField.Index, scortField.Type, scortField.Tag.Get("smurfs")) // .Tag.Get("Tag标签名") 可以获取指定的Tag标签信息
	}

	// 获取结构体所包含的方法
	fmt.Printf("\n############ 获取结构体所包含的方法 ############\n")
	fmt.Printf("当前结构体方法数量: %d\n", t.NumMethod()) // 获取当前结构体所包含的方法数量
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("method name:%s\n", t.Method(i).Name) // 通过Method()方法获取指定方法的Name信息
		fmt.Printf("method:%s\n", v.Method(i).Type())    // 通过Method()方法获取指定方法的Type()签名

		// 通过反射调用方法，传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}          // 定义一个用于存放参数的reflect.Value类型的切片
		if t.Method(i).Name == "SetStudent" { // 获取方法名，并判断此方法是不是需要参数的方法
			args = append(args, reflect.ValueOf("Smurfs")) // 将参数转换为reflect.Value类型并将其追加到args切片中
			args = append(args, reflect.ValueOf(99))
		}
		v.Method(i).Call(args) // 调用参数并传入参数。及时对应的方法没有设置形参，也需要传一个空的reflect.Value的切片进去
		fmt.Println()
	}
}
