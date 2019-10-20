package fiv

import (
	"fmt"
	"math/rand"
	"time"
)

// Fiv 设计一个程序，存储学员的信息: id 姓名 年龄 分数, 并且能够根据ID获取学员信息
func Fiv() {
	/*
		创建一个map
			map的键为int类型
			map的值也是一个map(这个map的键是string，值也是string)
	*/

	// rand.Seed方法可以让每次运行程序产生的随机数都不一样(前提是传入的整形参数是需要一直在变化的)
	rand.Seed(time.Now().UnixNano())

	var student_list = make(map[int]map[string]string, 10)

	for i := 0; i < 10; i++ {
		student_list[i] = make(map[string]string, 3)
		student_list[i]["姓名"] = fmt.Sprintf("Student_%d", i)
		student_list[i]["年龄"] = fmt.Sprintf("%d", rand.Intn(30))
		student_list[i]["分数"] = fmt.Sprintf("%d", rand.Intn(999))
	}

	// for k1, v1 := range student_list {
	// 	fmt.Println("学员ID: ", k1)
	// 	for k2, v2 := range v1 {
	// 		fmt.Printf("\t %s: %s\n", k2, v2)
	// 	}
	// }

	for i := 0; i < 10; i++ {
		fmt.Println(student_list[i])
	}
}
