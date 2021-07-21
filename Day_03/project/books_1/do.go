package books_1

import (
	"fmt"
	"os"
)

func book_obj() (name, writer string, price float32, publish bool) {
	for {
		fmt.Printf("书名: ")
		fmt.Scanln(&name)

		var status bool
		for _, v := range book_list {
			if (*v).Name == name {
				fmt.Println("书名重复了，请重新输入")
				status = true
				break
			}
		}
		if status == false {
			break
		}
	}
	fmt.Printf("作者姓名: ")
	fmt.Scanln(&writer)
	fmt.Printf("价格: ")
	fmt.Scanln(&price)
	fmt.Printf("是否上架: ")
	fmt.Scanln(&publish)

	return
}

func add() {
	name, writer, price, publish := book_obj()

	var book_res = &book{
		Name:    name,
		Writer:  writer,
		Price:   price,
		Publish: publish,
	}

	book_list = append(book_list, book_res)
	fmt.Println()
}

func set() {
Exit:
	for {
		var book_name string
		fmt.Print("请输入要修改书的名字: ")
		fmt.Scanln(&book_name)
		var status bool
		for _, v := range book_list {
			if (*v).Name == book_name {
				(*v).Name = ""
				name, writer, price, publish := book_obj()

				(*v).Name = name
				(*v).Writer = writer
				(*v).Price = price
				(*v).Publish = publish
				status = true
				fmt.Printf("修改完毕\n\n")
				break Exit
			}
		}
		if status == false {
			fmt.Println("没有这本书")
		}
	}
}

func sel() {
	for k, v := range book_list {
		fmt.Printf("########%d########\n", k)
		fmt.Printf("书名: 《%s》\n", (*v).Name)
		fmt.Printf("作者: %s\n", (*v).Writer)
		fmt.Printf("价格: %.2f元\n", (*v).Price)
		fmt.Printf("是否上架: %t\n", (*v).Publish)
	}
	fmt.Println()
}

func Do() {
	for {
		var do int

		fmt.Println("1. 添加书籍")
		fmt.Println("2. 修改书籍")
		fmt.Println("3. 查看书籍")
		fmt.Println("4. 退出")
		fmt.Printf("\n%s: ", "请输入您的操作")
		fmt.Scanln(&do)

		switch do {
		case 1:
			add()
		case 2:
			set()
		case 3:
			sel()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("输入错误")
			break
		}
	}
}
