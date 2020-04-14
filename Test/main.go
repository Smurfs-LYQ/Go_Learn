package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	res, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(res))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/index/", indexHandle)
	server.ListenAndServe()
}
