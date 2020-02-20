package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	str, err := ioutil.ReadFile("./name.txt")
	if err != nil {
		fmt.Println("打开name文件失败, err:", err)
		return
	}
	name_list := strings.Split(string(str), "\n")

	err = os.Chdir("/Volumes/Data/63107031")
	if err != nil {
		fmt.Println("进入文件夹失败, err:", err)
		return
	}
	fileInfoList, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(fileInfoList); i++ {
		wg.Add(1)
		filename := fileInfoList[i].Name()

		err = os.Rename(filename, name_list[i])
		if err != nil {
			fmt.Println(err)
			fmt.Println(filename, name_list[i])
		}
		wg.Done()

		wg.Wait()

	}
}
