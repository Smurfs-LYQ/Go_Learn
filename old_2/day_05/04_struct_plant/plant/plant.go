package plant

import (
	"Go_Learn/day_05/04_struct_plant/model"
	"fmt"
)

// Golang中的struct没有构造函数, 一般可以使用工厂模式来解决这个问题

func Plant_1() {
	/*
		// No.1
		// 初始化一个值类型的变量
		s := new(model.man) // 因为model中的man方法，方法名是小写的，所以导致外界不能访问
		// 创建这个实例
		s = model.NewMan("Smurfs的格格巫", 21)
		fmt.Println(s)
	*/

	// 创建并初始化一个实例
	s1 := model.NewMan("Smurfs的格格巫", 21)
	fmt.Println(s1)
}
