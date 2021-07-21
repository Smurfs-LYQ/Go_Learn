package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	// gResult 映射到从搜索拿到的结果文档
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UneacapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"conntent"`
	}

	// gResponse 包含顶级的文档
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func main() {
	uri := "http://ajax.googleaips.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	// 向 Google 发起搜索
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("Error: ", err)
		return
	}

	defer resp.Body.Close()

	// 将 JSON 相应解码到结构体
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("Error: ", err)
		return
	}

	fmt.Println(gr)
}
