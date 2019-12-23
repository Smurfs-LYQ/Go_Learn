package main

import "fmt"

// 数据类型
func main() {
	
	// 布尔值
	// var a bool
	// a = true
	// fmt.Println(a)

	// 整形
	// var a int = (1+1)*2 //算数表达式
	// var b int = 10<<2 //位操作表达式
	// fmt.Println(a)
	// fmt.Println(b)

	// 复数类型
	// var res1 complex64 = 3.1 + 5i
	// res2 := 3.1 + 6i
	// fmt.Println(res1)
	// fmt.Println(res2)
	// -----------------
	// var v = complex(2.1, 19) // 构建一个复数
	// a := real(v)
	// b := imag(v)
	// fmt.Println(v)
	// fmt.Println(a)	//返回复数实部
	// fmt.Println(b)  //返回复数虚部

	// 字符串
	// var a = "hello,world"

	// // 字符串切片输出，切片要求及格式与Python相同
	// b := a[0:5]
	// fmt.Println(b)

	// // 字符串转换为字节数组
	// c := []byte(a)
	// fmt.Println(c[6:11])

	// // 字符串转换为Unicode数组
	// d := []rune(a)
	// fmt.Println(d[6:11])

	// // 字符串拼接
	// e := "Hello"
	// f := "golang"
	// g := e + "," + f
	// fmt.Println(len(g))
	// fmt.Println(g)

	// for i := 0; i < len(g); i++ {
	// 	fmt.Println(g[i])
	// }

	// // range可以自动给每次循环的值添加一个下标，而且会自增
	// for i, v := range g{
	// 	fmt.Println(i,v)
	// }

	// 指针
	// var a *int
	// fmt.Printf("%v\n", a)

	// *a = 6
	// fmt.Printf("%v\n", *a)

	// var p *int
	// fmt.Printf("%v\n",p) //← 打印 nil

	// var i int //← 定义一个整形变量 i
	// p = &i //← 使得 p 指向 i， 获取 i 的地址
	// fmt.Printf("%v\n",p) //打印内存地址
 
	// *p = 6
	// fmt.Printf("%v\n",*p) //打印6
	// fmt.Printf("%v\n",i) //打印6

	// 1. 创建一个结构体，声明变量
	type User struct {
		name string
		age int
	}

	// 复制一个结构体并给他赋值
	andes := User {
		name: "Smurfs",
		age: 18,
	}

	// 将andes的内存地址赋值给指针p
	p := &andes
	fmt.Println(p)
	fmt.Println(p.name)
}