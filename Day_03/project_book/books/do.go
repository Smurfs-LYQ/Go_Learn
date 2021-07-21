package books

import "fmt"

func Book_Set() (name string, writer string, price float32, publish bool) {
	fmt.Printf("书名: ")
	fmt.Scanln(&name)
	fmt.Printf("作者姓名: ")
	fmt.Scanln(&writer)
	fmt.Printf("价格: ")
	fmt.Scanln(&price)
	fmt.Printf("是否上架: ")
	fmt.Scanln(&publish)

	return
}

func Add() {
	name, writer, price, publish := Book_Set()

	var book = book{
		name,
		writer,
		price,
		publish,
	}
	book_list = append(book_list, &book)
	fmt.Printf("添加成功\n\n")
}

func Set() {
	for k, v := range book_list {
		fmt.Printf("%d.《%s》\n", k, v.Name)
	}
	var book_id int
	for {
		fmt.Printf("\n请输入要修改哪本书: ")
		fmt.Scanln(&book_id)
		if len(book_list)-1 < book_id || book_id < 0 {
			fmt.Println("请输入正确数值!")
		} else {
			break
		}
	}

	name, writer, price, publish := Book_Set()
	(*book_list[book_id]).Name = name
	(*book_list[book_id]).Writer = writer
	(*book_list[book_id]).Price = price
	(*book_list[book_id]).Publish = publish
	fmt.Println("修改完毕", *book_list[book_id])
}

func Show() {
	for _, v := range book_list {
		fmt.Println("#####################")
		fmt.Printf("书名: %s\n", (*v).Name)
		fmt.Printf("作者: %s\n", (*v).Writer)
		fmt.Printf("价格: %.2f元\n", (*v).Price)
		fmt.Printf("是否上架: %t\n", (*v).Publish)
	}
	fmt.Println()
}
