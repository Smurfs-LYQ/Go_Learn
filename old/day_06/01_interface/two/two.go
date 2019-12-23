package two

import (
	"fmt"
	"math/rand"
	"sort"
)

/*
通过实现sort.Sort()接口中所包含的方法，就可以直接调用此接口进行排序。下面是一个通过实现sort.Sort()接口中所有方法从而对学生进行排序的实例
通过这样的方式就不用每次单独给一个类型排序了，可以通过多态的方式进行排序，可以有效的节省代码
*/

// 创建一个学生的结构体
type Student struct {
	Id   int
	Name string
	Age  int
}

// 给学生结构体取一个别名，并且将此结构体设置为切片类型的结构体
type StudentArray []Student

// 实现sort.Sort()接口中的方法
func (stu StudentArray) Len() int {
	return len(stu)
}

// 实现sort.Sort()接口中的方法
func (stu StudentArray) Less(i, j int) bool {
	return stu[i].Age < stu[j].Age
}

// 实现sort.Sort()接口中的方法
func (stu StudentArray) Swap(i, j int) {
	stu[i], stu[j] = stu[j], stu[i]
}

func Two() {
	// 创建一个学生数组的结构体实例
	var stu_array StudentArray

	// 创建学生并赋值
	for i := 0; i < 5; i++ {
		var stu = Student{
			Id:   rand.Int(),
			Name: fmt.Sprintf("stu_%d", rand.Intn(100)),
			Age:  rand.Intn(20),
		}

		stu_array = append(stu_array, stu)
	}

	// 打印出创建的学生列表
	for _, v := range stu_array {
		fmt.Println(v)
	}

	fmt.Printf("\n\n")
	// 通过sort.Sort()接口进行排序
	sort.Sort(stu_array)

	// 打印出排序之后的学生列表
	for _, v := range stu_array {
		fmt.Println(v)
	}
}
