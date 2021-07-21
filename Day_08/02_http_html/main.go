package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		panic(fmt.Sprintln("页面加载失败, err:", err))
	}

	html := string(data)

	if rand.Seed(time.Now().UnixNano()); rand.Intn(2) == 1 {
		html = strings.Replace(html, "{tag_1}", "<li> one </li>", 1)
	} else {
		html = strings.Replace(html, "{tag_1}", "<li> zero </li>", 1)
	}
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
