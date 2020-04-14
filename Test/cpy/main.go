package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	json := strings.NewReader(`"CPYYXZKCCX~{'NSRSBH':'1110115MA00GMBB48','FJH':'1','SPBM':'1070101070100000000'}"`)
	// req, _ := http.NewRequest("GET", "http://127.0.0.1:8080/zzs_kpfw_ARM9/pagemanage", json)
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8080/index/", json)

	req.Header.Add("Accept", `*/*`)
	req.Header.Add("Content-Type", `application/x-www-form-urlencoded`)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
