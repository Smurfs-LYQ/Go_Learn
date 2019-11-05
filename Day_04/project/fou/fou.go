package fou

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Cat 实现一个cat命令
func Cat() {
	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Println("请传入要打印的参数")
		return
	}

	for _, v := range args {
		file, err := ioutil.ReadFile(v)
		if err != nil {
			fmt.Println("读取文件失败, 失败原因: ", err)
			continue
		}
		fmt.Printf("#########%s#########\n", v)
		fmt.Println(string(file))
	}
}
