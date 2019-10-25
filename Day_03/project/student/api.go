package student

import "fmt"

type student struct {
	ID    int
	Name  string
	Age   int
	Class string
}

var student_list = make([]*student, 0, 10)

func student_message() (name string, age int, class string) {
	fmt.Printf("请输入学生姓名: ")
	fmt.Scanln(&name)
	fmt.Printf("请输入学生年龄: ")
	fmt.Scanln(&age)
	fmt.Printf("请输入学生班级: ")
	fmt.Scanln(&class)

	return
}

func (s student) Add() {
	for _, v := range student_list {
		if v.ID >= s.ID {
			s.ID = v.ID + 1
		}
	}
	s.Name, s.Age, s.Class = student_message()
	student_list = append(student_list, &s)
	fmt.Printf("%s\n\n", "添加成功")
}

func Set() {
Begin:
	fmt.Printf("请输入要修改的学生姓名: ")
	var name string
	fmt.Scanln(&name)
	var status bool
	for _, v := range student_list {
		if v.Name == name {
			// status = true
			v.Set()
			fmt.Printf("修改完成\n\n")
			return
		}
	}
	if !status {
		fmt.Println("没有这个学生")
		goto Begin
	}
}

func (s *student) Set() {
	s.Name, s.Age, s.Class = student_message()
}

func Del() {
	fmt.Printf("请输入要删除的学生姓名: ")
	var name string
	fmt.Scanln(&name)
	var key int
	for k, v := range student_list {
		if v.Name == name {
			key = k
			break
		}
	}
	if key == 0 {
		if student_list[0].Name != name {
			fmt.Printf("没有这个学生\n\n")
			return
		}
		student_list = append(student_list[1:])
		fmt.Printf("已删除\n\n")
		return
	}
	student_list = append(student_list[:key], student_list[key+1:]...)
	fmt.Printf("已删除\n\n")
	return
}

func Show() {
	for _, v := range student_list {
		fmt.Printf("########学生ID: %d########\n", v.ID)
		fmt.Printf("学生姓名: %s\n", v.Name)
		fmt.Printf("学生年龄: %d\n", v.Age)
		fmt.Printf("学生班级: %s\n", v.Class)
	}
	fmt.Println()
}
