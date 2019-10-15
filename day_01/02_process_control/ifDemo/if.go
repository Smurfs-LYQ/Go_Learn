package ifDemo

import "fmt"

// One if示例函数
func One() {
	age := 21

	// if 流程控制基础写法
	if age < 18 {
		fmt.Println("这玩意是童工")
	} else if age > 18 && age < 60 {
		fmt.Println("兄弟你马上就要退休了")
	} else {
		fmt.Println("终于退休了")
	}

	// if 流程控制特殊写法
	if sex := "Man"; sex == "Man" {
		fmt.Println("男的")
	} else {
		fmt.Println("女的")
	}
}
