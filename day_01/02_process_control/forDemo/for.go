package forDemo

import "fmt"

// One for循环示例
func One() {
	// for 循环基础写法
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// for循环的初始语句可以被忽略，但是初始语句后的分号必须写
	a := 10
	for ; a < 20; a++ {
		fmt.Printf("%d ", a)
	}
	fmt.Println()

	// for循环的初始语句和结束语句都可以省略
	b := 20
	for b < 30 {
		fmt.Printf("%d ", b)
		b++
	}
	fmt.Println()

	// for循环break + Label标签使用案例
	for i := 0; i < 5; i++ {
		// 设置退出Label标签
	exitTag:
		for o := 0; o < i; o++ {
			if i == 3 {
				fmt.Printf("这里是跳过不输入的地点")
				// break到退出标签
				break exitTag
			}
			fmt.Printf("%d-%d ", o, i)
		}
		fmt.Println()
	}

	// for循环continue + Label标签使用案例
	for i := 0; i < 5; i++ {
		// 设置退出Label标签
	exitTag_1:
		for o := 0; o <= i; o++ {
			if o == 3 && i == 3 {
				fmt.Printf("这里是跳过不输入的地点")
				// continue到退出标签
				continue exitTag_1
			}
			fmt.Printf("%d-%d ", o, i)
		}
		fmt.Println()
	}
}
