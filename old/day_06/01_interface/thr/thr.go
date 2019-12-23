package thr

import "fmt"

/*
接口嵌套
	一个接口可以嵌套在另外的接口，效果跟结构体的嵌套类似，如下所示：
		type ReadWrite interface {
			Read(url) string
			Write(url) bool
		}

		type Lock interface {
			Lock()
			Unlock()
		}

		type File interface {
			ReadWrite
			Lock
			Close()
		}
*/

// 定义一个读的接口
type Reader interface {
	Read()
}

// 定义一个写的接口
type Writer interface {
	Write()
}

// 定义一个嵌套和写的接口
type ReadWriter interface {
	Reader // 嵌套读的接口那么他就包含了 Read() 方法
	Writer // 嵌套写的接口那么他就包含了 Write() 方法
}

// 定义一个结构体
type File struct {
	Name string
}

// 实现Reader接口的方法
func (this File) Read() {
	fmt.Println(this.Name, "read data")
}

// 实现Writer接口的方法
func (this File) Write() {
	fmt.Println(this.Name, "Write data")
}

// 定义一个方法，参数为同时实现了Reader和Writer中所有方法的结构体变量
func Test(this ReadWriter) {
	this.Read()
	this.Write()
}

func Thr() {
	// 创建一个结构体变量
	var file File = File{
		Name: "Smurfs",
	}

	// 调用Test方法，并传入同时实现了Reader和Writer中所有方法的结构体变量
	Test(file)
}
