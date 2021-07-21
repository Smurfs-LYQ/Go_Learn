package General

import "fmt"

// One 通用占位符
func One() {
	// %v 以默认的方式打印变量的值
	fmt.Printf("%v %v %v\n", "%v什么类型都可以打印，比如:", 123, true)
	// %T 打印变量的类型
	fmt.Printf("%T %T %T\n", "test", 123, true)
	// %% 字面上的百分号，并非占位符
	fmt.Printf("%s %%\n", "%%的意思就是")
}
