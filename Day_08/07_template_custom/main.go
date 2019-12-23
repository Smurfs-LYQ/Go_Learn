package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 读取页面
	htmlByte, err := ioutil.ReadFile("./index.html")
	if err != nil {
		fmt.Println("打开文件失败, err:", err)
		return
	}

	// 在解析模板文件之前添加自定义方法
	// 1. 自定义一个函数
	HelloFunc := func(str string) (string, error) {
		return "Hello " + str, nil
	}
	// 2. 把自定义的函数告诉模板文件
	// template.New("对象名") 创建一个模板对象
	// .Funcs(template.FuncMap{"传入模板中的函数名": 执行操作的函数名}) 给模板追加自定义函数
	// .Parse 解析的页面
	t, err := template.New("index").Funcs(template.FuncMap{"Hello": HelloFunc}).Parse(string(htmlByte))
	if err != nil {
		fmt.Println("页面加载失败, err:", err)
		return
	}

	MotoList := map[int]string{
		1: "Kawasaki",
		2: "Yamaha",
		3: "Honda",
		4: "Suzuki",
	}

	t.Execute(w, MotoList)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
