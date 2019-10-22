package scan

import "fmt"

func One() {
	var (
		name string
		age  int
		sex  string
	)

	// Scan 从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
	/*
		因为fmt.Scan()中传入了3个参数，所以这里截取了3文本后会自动停止
	*/
	fmt.Scan(&name, &age, &sex)
	fmt.Println(name, age, sex)
}
