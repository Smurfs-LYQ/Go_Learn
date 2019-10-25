package books_2

import (
	"fmt"
	"os"
)

func Do() {
	for {
		var do int

		fmt.Println("1. 添加书籍")
		fmt.Println("2. 修改书籍")
		fmt.Println("3. 查看书籍")
		fmt.Println("4. 退出")
		fmt.Printf("\n%s: ", "请输入您的操作")
		fmt.Scanln(&do)

		var obj book

		switch do {
		case 1:
			obj.add()
		case 2:
			set()
		case 3:
			show()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("输入错误")
			break
		}
	}
}
