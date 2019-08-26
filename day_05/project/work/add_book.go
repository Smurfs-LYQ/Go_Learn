package work

import (
	"fmt"
	"os"
	"strings"
)

// 添加书籍

func Add_Book() {
	// fmt.Println("添加书籍")
	list := []string{"书名", "副本数", "作者", "出版日期"}
	var res_list []string
	for _, v := range list {
		var res string
	Label1:
		fmt.Printf("请输入%s: ", v)
		fmt.Scanf("%s\n", &res)
		// 如果输入为空返回重新输入
		if res == "" {
			goto Label1
		}
		res_list = append(res_list, res)
	}

	fmt.Println(res_list)

	fileName := "../data/books.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0766)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	defer file.Close()

	file.Seek(0, 2) // 最后增加
	file.WriteString(strings.Join(res_list, "#_#") + "\r\n")
}
