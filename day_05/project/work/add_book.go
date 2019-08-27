package work

// import "fmt"
import (
	"fmt"
)

// import "fmt"

// 添加书籍
// 书籍结构体
type Books struct {
	Name   string `json:"name"`   // 书名
	Count  int    `json:"count"`  // 副本数
	Writer string `json:"writer"` // 作者
	Time   string `json:"time"`   // 出版日期
}

var one Books = Books{
	Name:   "test1",
	Count:  12,
	Writer: "Smurfs",
	Time:   "2019/3/2",
}
var two Books = Books{
	Name:   "test2",
	Count:  42,
	Writer: "Smurfs",
	Time:   "2019/8/2",
}
var thr Books = Books{
	Name:   "test4",
	Count:  5,
	Writer: "test",
	Time:   "2019/3/1",
}
var Books_list = []*Books{&one, &two, &thr}

func Add_Book() {
	fmt.Println("添加书籍")
	var name string
	fmt.Printf("请输入书名: ")
	fmt.Scanf("%s\n", &name)
	var count int
	fmt.Printf("请输入副本数: ")
	fmt.Scanf("%d\n", &count)
	var writer string
	fmt.Printf("请输入作者: ")
	fmt.Scanf("%s\n", &writer)
	var time string
	fmt.Printf("请输入出版日期: ")
	fmt.Scanf("%s\n", &time)

	var book = &Books{
		Name:   name,
		Count:  count,
		Writer: writer,
		Time:   time,
	}

	Books_list = append(Books_list, book)
	fmt.Println(*book)
}
