package student_2

import "fmt"

type Student struct {
	ID    int
	Name  string
	Age   int
	Class string
}

func (s Student) NewStudent(id int, name string, age int, class string) *Student {
	s.ID = id
	s.Name = name
	s.Age = age
	s.Class = class
	return &s
}

func (s *Student) SetStudent(id int, name string, age int, class string) {
	s.ID = id
	s.Name = name
	s.Age = age
	s.Class = class
}

type Student_Manage struct {
	List []*Student
}

func (s *Student_Manage) Add(stu *Student) {
	for _, v := range s.List {
		if v.Name == stu.Name {
			fmt.Println("名字已被占用")
			return
		}
	}
	s.List = append(s.List, stu)
	fmt.Println("添加完成")
}

func (s *Student_Manage) Set() {
	fmt.Print("请输入要修改学生的姓名: ")
	var name string
	fmt.Scanln(&name)

	for _, v := range s.List {
		if (*v).Name == name {
			id, name, age, class := Input()
			v.ID = id
			v.Name = name
			v.Age = age
			v.Class = class
			fmt.Println("修改成功")
		}
	}
	fmt.Println("没有这个学生")
}

func (s *Student_Manage) Show() {
	for _, v := range s.List {
		fmt.Println("##########################")
		fmt.Printf("%s%d\n", "学生ID:", v.ID)
		fmt.Printf("%s%s\n", "学生姓名:", v.Name)
		fmt.Printf("%s%d\n", "学生年龄:", v.Age)
		fmt.Printf("%s%s\n", "学生班级:", v.Class)
	}
}

func (s *Student_Manage) Del() {
	fmt.Print("请输入要删除学生的姓名: ")
	var name string
	fmt.Scanln(&name)

	for k, v := range s.List {
		if (*v).Name == name {
			s.List = append(s.List[:k], s.List[k+1:]...)
			fmt.Println("删除成功")
			return
		}
	}
	fmt.Println("删除失败，没有这个学生")
}

func Input() (id int, name string, age int, class string) {
	fmt.Printf("请输入学生学号: ")
	fmt.Scanln(&id)
	fmt.Printf("请输入学生姓名: ")
	fmt.Scanln(&name)
	fmt.Printf("请输入学生年龄: ")
	fmt.Scanln(&age)
	fmt.Printf("请输入学生班级: ")
	fmt.Scanln(&class)

	return
}
