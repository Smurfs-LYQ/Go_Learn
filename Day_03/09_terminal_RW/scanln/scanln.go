package scanln

import "fmt"

func One() {
	var (
		name string
		age  int
		sex  string
	)
	// Scanln其实和Scan是差不多的，但是Scanln只能用空格进行分割，检测到换行就自动停止了
	fmt.Scanln(&name, &age, &sex)
	fmt.Printf("%q, %d, %q\n", name, age, sex)
}
