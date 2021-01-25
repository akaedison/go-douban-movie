package parse

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type Page struct {
	Page int
	Url  string
}

func GetPages(url string) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1;WOW64) AppleWebKit/537.36 (KHTML,like GeCKO) Chrome/45.0.2454.85 Safari/537.36 115Broswer/6.0.3")
	req.Header.Set("Referer", "https://movie.douban.com/")
	req.Header.Set("Connection", "keep-alive")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		panic("读取失败")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("连接失败")
	}
	//body, err := ioutil.ReadAll(res.Body)

	//reader := strings.NewReader(string(body))
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	html, _ := doc.Html()
	fmt.Println(html)
}
