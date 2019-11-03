package thr

import (
	"fmt"
	"io/ioutil"
)

// Thr ioutil.WriteFile写入文件实例
func Thr() {
	err := ioutil.WriteFile("../xx.txt", []byte("要写入的内容"), 0755)
	if err != nil {
		fmt.Println("写入失败，错误信息: ", err)
		return
	}
}
