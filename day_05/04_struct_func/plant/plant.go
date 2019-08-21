package plant

import (
	"Go_Learn/day_05/04_struct_func/model"
	"fmt"
)

// Golang中的struct没有构造函数, 一般可以使用工厂模式来解决这个问题

func Plant_1() {
	/*
		// No.1
		// 初始化一个值类型的变量
		s := new(model.Man)
		// 创建这个实例
		s = model.NewMan("Smurfs的格格巫", 21)
		fmt.Println(s)
	*/

	// 创建并初始化一个实例
	s := model.NewMan("Smurfs的格格巫", 21)
	fmt.Println(s)
}
