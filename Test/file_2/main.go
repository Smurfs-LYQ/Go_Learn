package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var pwd = "/Volumes/Data/go/"

func main() {
	str, err := ioutil.ReadFile("./a.txt")
	if err != nil {
		fmt.Println("打开name文件失败, err:", err)
		return
	}
	nameList := strings.Split(string(str), "\n")

	err = os.Chdir(pwd)
	if err != nil {
		panic(fmt.Sprintf("进入文件夹失败, err:%v\n", err))
	}

	fileInfoList, err := ioutil.ReadDir("./")
	if err != nil {
		panic(fmt.Sprintf("获取目录信息失败, err:%v\n", err))
	}

	for k, v := range fileInfoList {
		file := fmt.Sprintf("%s%s/", pwd, v.Name())

		do(file, nameList[k])

		err := os.Rename(v.Name(), nameList[k])
		if err != nil {
			fmt.Printf("%s 改名失败\n", v.Name())
		}
	}
}

func do(file string, name string) {
	err := os.Chdir(file)
	if err != nil {
		panic(fmt.Sprintf("%s, 进入目录失败, err:%v\n", file, err))
	}
	defer os.Chdir("../")

	fileInfoList, err := ioutil.ReadDir("./")
	if err != nil {
		panic(fmt.Sprintf("%s, 获取目录信息失败, err:%v\n", file, err))
	}

	for _, v := range fileInfoList {
		res := strings.Split(v.Name(), ".")
		if res[1] == "flv" {
			oldName := fmt.Sprintf("./%s", v.Name())
			newName := fmt.Sprintf("./%s.flv", name)

			err := os.Rename(oldName, newName)
			if err != nil {
				fmt.Printf("%s%s 改名失败\n", file, v.Name())
			}
		} else {
			err := os.Remove(v.Name())
			if err != nil {
				fmt.Printf("%s%s 删除失败\n", file, v.Name())
			}
		}
	}
}
