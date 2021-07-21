package student_1

import (
	"fmt"
	"os"
)

/*
学生管理系统
	字段: 姓名 年龄 ID 班级
	1. 增加学生
	2. 修改学生
	3. 删除学生
	4. 展示学生
	要求用结构体加方法的形式
*/

func Do() {
	for {
		var do int

		fmt.Println("1. 增加学生")
		fmt.Println("2. 修改学生")
		fmt.Println("3. 删除学生")
		fmt.Println("4. 展示学生")
		fmt.Println("5. 退出")

		fmt.Printf("请输入操作选项: ")
		fmt.Scanln(&do)

		var student_obj student

		switch do {
		case 1:
			student_obj.Add()
		case 2:
			Set()
		case 3:
			Del()
		case 4:
			Show()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("输入错误")
			break
		}
	}
}
