package gotoDemo

import "fmt"

// One goto示例
func One() {

	for i := 0; i < 10; i++ {
		if i > 6 {
			// goto到退出标签
			goto exitTag
		}
		fmt.Printf("%d ", i)
	}
	// 设置退出labal标签
exitTag:
	fmt.Printf("\n%s\n", "到6就行了")
}
