package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var wg sync.WaitGroup

func do(file string) {
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

	newName := fmt.Sprintf("../%s", fileInfoList[0].Name())
	os.Rename(fileInfoList[0].Name(), newName)
	os.Remove(file)
}

func main() {
	err := os.Chdir("/Volumes/Data/63107031")
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
		filePath := fmt.Sprintf("/Volumes/Data/63107031/%s/", fileInfoList[i].Name())
		err := os.Chdir(filePath)
		if err != nil {
			fmt.Println("进入文件夹失败, err:", err)
			return
		}

		fileList, err := ioutil.ReadDir("./")
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, v := range fileList {
			wg.Add(1)
			go do(fmt.Sprintf("%s%s", filePath, v.Name()))
			wg.Wait()
		}
		fmt.Println()

		if err != nil {
			fmt.Println("退出文件夹失败, err:", err)
			return
		}
	}
}
