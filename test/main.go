package main

import (
	"fmt"
	"strings"
)

func urlFormat(url string) string {
	judge := strings.HasPrefix(url, "https://")
	if !judge {
		url = "https://" + url
	}
	return url
}

func pathFormat(path string) string {
	judge := strings.HasSuffix(path, "/")
	if !judge {
		path = path + "/"
	}
	return path
}

func main() {
	var url string
	var path string

	fmt.Scanf("%s %s", &url, &path)

	newurl := urlFormat(url)
	newpath := pathFormat(path)

	fmt.Println(newurl)
	fmt.Println(newpath)
}
