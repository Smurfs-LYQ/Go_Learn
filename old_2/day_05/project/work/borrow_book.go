package work

import "fmt"

func Borrow_Book() {
	// fmt.Println("借书")
	var stu_id string
	var book_name string
	var book_num int
Label1:
	fmt.Printf("学生ID: ")
	fmt.Scanf("%s\n", &stu_id)

	var stu *Student
	str_status := false
	for i := 0; i < len(Student_List); i++ {
		stu = Student_List[i]
		if stu.ID == stu_id {
			str_status = true
			break
		}
	}
	if str_status == false {
		fmt.Println("没有这个学生")
		goto Label1
	}
Label2:
	fmt.Printf("想看什么书: ")
	fmt.Scanf("%s\n", &book_name)

	var book *Books
	book_status := false
	for i := 0; i < len(Books_list); i++ {
		book = Books_list[i]
		if book.Name == book_name {
			book_status = true
			break
		}
	}
	if book_status != true {
		fmt.Println("没有这本书")
		goto Label2
	}

Label3:
	fmt.Printf("想借多少本: ")
	fmt.Scanf("%d\n", &book_num)

	if book.Count < book_num {
		if book.Count == 0 {
			fmt.Println("这本书没有库存了")
			return
		}
		fmt.Printf("没呢么多，只有%d本\n", book.Count)
		goto Label3
	}

	fmt.Println(*book)

	// 操作用户信息和书的信息
	var stu_book Stu_Books = Stu_Books{
		user_book: book,
		num:       book_num,
	}
	book.Count -= book_num

	stu.Book = append(stu.Book, &stu_book)
}
