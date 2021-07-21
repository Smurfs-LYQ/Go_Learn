package work

import "fmt"

func Search_Book() {
	// fmt.Println("搜索书籍")
	fmt.Println()
	if len(Books_list) <= 0 {
		fmt.Println("没书")
	}

	fmt.Printf("请输入作者名: ")
	var name string
	fmt.Scanf("%s\n", &name)

	fmt.Println()
	status := false
	for _, v := range Books_list {
		one := *v
		if one.Writer == name {
			status = true
			fmt.Println(*v)
		}
	}

	if status == false {
		fmt.Println("没有这个作者的书")
	}
}
