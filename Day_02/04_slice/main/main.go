package main

import "fmt"

func main() {
	// 声明一个切片
	var one []int
	fmt.Printf("%T %d\n", one, one) // 因为是int类型的slice所以可以用%d来占位

	// 声明并初始化一个切片
	var two = []int{1, 2, 3}
	fmt.Printf("%T %d %d\n", two, two, two[1])

	// 基于数组定义切片
	var arr = [5]string{"one", "two", "thr", "fou", "fiv"}
	fmt.Printf("%T %s\n", arr, arr)
	// 对数组进行切片
	var thr = arr[1:3]
	fmt.Printf("%T %s\n", thr, thr)

	// 检测切片的大小 (目前元素的数量)
	fmt.Println(len(thr))
	// 检测切片的容量 (底层数组最大能放多少元素)))
	fmt.Println(cap(thr)) // 切片的容量是从切片开始切数组的位置一直到数组的最后。因为arr的总长度是5，而thr的内存地址是根据arr索引为1的位置记录的，所以thr的容量为5-1=4

	// 切片元素增加，切片容量的变化规律
	/*
		切片的扩容策略: 切片的容量不够时，切片会自动扩充容量，扩充的大小为扩充前容量的一倍
	*/
	var fou = []int{}
	fmt.Printf("内容:%d 长度:%d 容量:%d 指针:%p\n", fou, len(fou), cap(fou), fou)

	fou = append(fou, 1)
	fmt.Printf("内容:%d 长度:%d 容量:%d 指针:%p\n", fou, len(fou), cap(fou), fou)

	fou = append(fou, 1)
	fmt.Printf("内容:%d 长度:%d 容量:%d 指针:%p\n", fou, len(fou), cap(fou), fou)

	fou = append(fou, 1)
	fmt.Printf("内容:%d 长度:%d 容量:%d 指针:%p\n", fou, len(fou), cap(fou), fou)

	fou = append(fou, 1)
	fmt.Printf("内容:%d 长度:%d 容量:%d 指针:%p\n", fou, len(fou), cap(fou), fou)

	fou = append(fou, 1)
	fmt.Printf("内容:%d 长度:%d 容量:%d 指针:%p\n", fou, len(fou), cap(fou), fou)

	// copy函数
	// 声明一个空切片
	var fiv []int // 这里虽然声明了，但是还没有申请内存
	// 通过make给引用类型分配内存
	fiv = make([]int, 3, 3) // 申请一块内存用于存放[]int类型切片，长度为 3，容量也为 3
	copy(fiv, two)
	fmt.Println(two, fiv)

	// 从切片中删除重复的元素"2"
	var six = []int{1, 2, 2, 3, 4, 5}
	fmt.Println(six)
	// 因为切片是引用类型，所以修改的就是其本身，使用append将需要保留的元素提取出来重新保存到该变量中
	six = append(six[:2], six[3:]...) // 这里的...操作的意思是将 six[3:] 中的元素挨个拆开，然后挨个放入six中
	fmt.Println(six)

	var a = []int{1, 2, 3, 4, 5, 6, 7}
	b := a[2:5]
	fmt.Println(a[2:5])
	c := b[:5]
	fmt.Println(c)
}
