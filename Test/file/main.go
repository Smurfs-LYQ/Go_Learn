package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func del(str1 string) {
	os.Chdir(str1)
	os.Remove("./danmaku.xml")
	os.Remove("./entry.json")
	os.Remove("./80/index.json")
	wg.Done()
}

func main() {
	str, err := ioutil.ReadFile("./name.txt")
	if err != nil {
		fmt.Println("打开name文件失败, err:", err)
		return
	}
	name_list := strings.Split(string(str), "\n")

	for _, v := range name_list {
		fmt.Println(v)
	}

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
		// err := os.Rename(fileInfoList[i].Name(), name_list[i])
		// if err != nil {
		// 	fmt.Println(err)
		// 	fmt.Println(fileInfoList[i].Name(), name_list[i])
		// }
		fmt.Println(fileInfoList[i].Name(), name_list[i])

		go del(fmt.Sprintf("/Volumes/Data/63107031/%s/", name_list[i]))

		wg.Wait()

	}
}
