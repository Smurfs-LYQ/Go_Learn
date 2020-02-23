package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func do(file, name string) {
	defer wg.Done()

	err := os.Chdir(file)
	if err != nil {
		fmt.Printf("进入文件夹%s失败, err:%v\n", file, err)
		return
	}
	fileInfoList, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(fileInfoList); i++ {
		if strings.Index(fileInfoList[i].Name(), "xml") > 0 || strings.Index(fileInfoList[i].Name(), "info") > 0 {
			os.Remove(fileInfoList[i].Name())
			continue
		}
		os.Rename(fileInfoList[i].Name(), name)
	}
}

func main() {
	str, err := ioutil.ReadFile("./file/name.txt")
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

		go do(fmt.Sprintf("/Volumes/Data/63107031/%s", fileInfoList[i].Name()), name_list[i])
		wg.Wait()
	}
}
