package student_2

import (
	"fmt"
	"os"
)

/*
学生管理系统
	字段: 姓名 年龄 学号 班级
	1. 增加学生
	2. 修改学生
	3. 删除学生
	4. 展示学生
	要求用结构体加方法的形式
*/

func Do() {
	var stu Student
	var student_man Student_Manage

	for {
		var do int

		fmt.Println("1. 增加学生")
		fmt.Println("2. 修改学生")
		fmt.Println("3. 删除学生")
		fmt.Println("4. 展示学生")
		fmt.Println("5. 退出")

		fmt.Printf("请输入操作选项: ")
		fmt.Scanln(&do)

		switch do {
		case 1:
			id, name, age, class := Input()
			res := stu.NewStudent(id, name, age, class)
			student_man.Add(res)
		case 2:
			student_man.Set()
		case 3:
			student_man.Del()
		case 4:
			student_man.Show()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("输入错误")
			break
		}
	}
}
