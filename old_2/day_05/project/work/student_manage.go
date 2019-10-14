package work

import "fmt"

type Student struct {
	Name  string
	Class string
	ID    string
	Sex   string
	Book  []*Stu_Books
}

type Stu_Books struct {
	user_book *Books
	num       int
}

var stu1 Student = Student{
	Name:  "one",
	Class: "9",
	ID:    "1",
	Sex:   "男",
}
var stu2 Student = Student{
	Name:  "two",
	Class: "9",
	ID:    "2",
	Sex:   "女",
}

var Student_List = []*Student{&stu1, &stu2}

func Add_Student() {
	var name string
	var class string
	var id string
	var sex string
	fmt.Printf("请输入姓名: ")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("请输入年级: ")
	fmt.Scanf("%s\n", &class)
	fmt.Printf("请输入身份证: ")
	fmt.Scanf("%s\n", &id)
	fmt.Printf("请输入性别: ")
	fmt.Scanf("%s\n", &sex)

	var stu Student = Student{
		Name:  name,
		Class: class,
		ID:    id,
		Sex:   sex,
	}
	Student_List = append(Student_List, &stu)
}

func Student_Manage() {
	// fmt.Println("学生信息管理")
	fmt.Println()
	fmt.Println("1. 添加学生")
	fmt.Println("2. 查看学生")
	fmt.Println()
	fmt.Printf("Next: ")
	var res int
	fmt.Scanf("%d\n", &res)
	fmt.Println()

	switch res {
	case 1:
		Add_Student()
	case 2:
		for _, v := range Student_List {
			// fmt.Println(*v)
			fmt.Println("姓名: ", v.Name)
			fmt.Println("年级：", v.Class)
			fmt.Println("I D: ", v.ID)
			fmt.Println("性别: ", v.Sex)
			for _, v := range v.Book {
				var book = *v.user_book
				fmt.Println("所借书籍: ")
				fmt.Printf("\t书名: %s\n", book.Name)
				fmt.Printf("\t作者: %s\n", book.Writer)
				fmt.Printf("\t出版日期: %s\n", book.Time)
				fmt.Printf("\t数量: %d\n", v.num)
				fmt.Println("------------------")
			}
			fmt.Println("##################")
		}
	default:
		return
	}

}
