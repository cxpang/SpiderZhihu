package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func gethome()  {
	fmt.Println("进入知乎我的首页")
	buildheader()
}
func buildheader()  {
	//第一个指定的url
	url:="https://www.zhihu.com/people/xnew/following"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	//req.Header.Set("accept-encoding", "gzip, deflate, br")
	req.Header.Set("accept-language", "accept-language")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("upgrade-insecure-requests", "1")
	resp, err := client.Do(req)
	if err!=nil {

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//找到所有的json格式文件
	ptnContentRough := regexp.MustCompile(`{"initialState":(.*)}`)
	match := ptnContentRough.FindStringSubmatch(string(body))

	fmt.Println(match)

}