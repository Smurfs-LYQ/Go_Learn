package main

import (
	"Go_Learn/day_05/project/work"
	"fmt"
)

/*
实现一个图书管理系统，具有以下功能
	* 书籍录入功能，书籍信息包括书名、副本数、作者、出版日期
	* 书籍查询功能，按照书名、作者、出版日期等条件检索
	* 学生信息管理功能，管理没个学生的姓名、年级、身份证、性别、借了什么书等信息
	* 借书功能，学生可以查询想要的书籍，进行借出
	* 书籍管理功能，可以看到每种书被哪些人借出了
*/

func main() {
	color := []int{31, 32, 33, 34, 35}
	for k, v := range []string{"书籍录入", "书籍查询", "学生信息管理", "借书", "书籍管理"} {
		fmt.Printf("%c[0;40;%dm%d. %s%c[0m\n", 0x1B, color[k], k, v, 0x1B)
	}

	fmt.Printf("\n%c[0;36m%s%c[0m", 0x1B, "要做什么: ", 0x1B)
	var one int
	fmt.Scanf("%d\r\n", &one)

	switch one {
	case 0:
		work.Add_Book()
	case 1:
		work.Search_Book()
	case 2:
		work.Student_Manage()
	case 3:
		work.Borrow_Book()
	case 4:
		work.Book_Manage()
	}
}
