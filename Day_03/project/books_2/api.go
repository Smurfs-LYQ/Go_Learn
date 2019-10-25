package books_2

import "fmt"

type book struct {
	Name    string
	Writer  string
	Price   float32
	Publish bool
}

var book_list = make([]*book, 0, 10)

func find_book(name string) bool {
	for _, v := range book_list {
		if (*v).Name == name {
			return true
		}
	}
	return false
}

func book_obj(do, book_name string) (name, writer string, price float32, publish bool) {
	for {
		fmt.Printf("书名: ")
		fmt.Scanln(&name)

		if do == "set" {
			if find_book(name) && name != book_name {
				fmt.Println("书名重复，请重新输入")
			} else {
				break
			}
		} else {
			if find_book(name) {
				fmt.Println("书名重复，请重新输入")
			} else {
				break
			}
		}
	}

	fmt.Printf("作者姓名: ")
	fmt.Scanln(&writer)
	fmt.Printf("价格: ")
	fmt.Scanln(&price)
	fmt.Printf("是否上架: ")
	fmt.Scanln(&publish)
	fmt.Println()

	return
}

func set() {
Begin:
	var book_name string
	fmt.Printf("请输入要修改的书名: ")
	fmt.Scanln(&book_name)

	if find_book(book_name) {
		// set(book_name)
		for _, v := range book_list {
			if (*v).Name == book_name {
				(*v).set()
			}
		}
	} else {
		fmt.Println("没有这本书")
		goto Begin
	}
}

func (b *book) add() {
	b.Name, b.Writer, b.Price, b.Publish = book_obj("add", "")

	book_list = append(book_list, b)
}

func (b *book) set() {
	b.Name, b.Writer, b.Price, b.Publish = book_obj("set", b.Name)

	fmt.Println("修改成功")
}

func show() {
	for k, v := range book_list {
		fmt.Printf("########%d########\n", k)
		fmt.Printf("书名: 《%s》\n", (*v).Name)
		fmt.Printf("作者: %s\n", (*v).Writer)
		fmt.Printf("价格: %.2f元\n", (*v).Price)
		fmt.Printf("是否上架: %t\n", (*v).Publish)
	}
	fmt.Println()
}
