package scanf

import "fmt"

func One() {
	var (
		name string
		age  int
		sex  string
	)
	/*
		Scanf跟Scan的区别在于，Scanf获取终端输入时必须严格按照设定的格式来填写，例如下面这个实例，在终端中需要填写的格式如下
		name:Smurfs age:21 sex:男
	*/
	fmt.Scanf("name:%s age:%d sex:%s\n", &name, &age, &sex)
	fmt.Println(name, age, sex)
}
