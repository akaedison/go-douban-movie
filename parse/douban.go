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

func getDoc(url string) *goquery.Document {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1;WOW64) AppleWebKit/537.36 (KHTML,like GeCKO) Chrome/45.0.2454.85 Safari/537.36 115Broswer/6.0.3")
	//req.Header.Set("Referer", "https://movie.douban.com/")
	req.Header.Set("Connection", "keep-alive")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		panic("读取失败")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("连接失败")
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	return doc
}

func getDouban(url string) {
	for i := 0; i < 10; i++ {
		url := fmt.Sprintf("%v?start=%v", url, i*25)
		doc := getDoc(url)
		doc.Find(".grid_view > li > .item").Each(func(i int, selection *goquery.Selection) {
			id := selection.Find(".pic em")
			fmt.Println(id)
		})
	}
}
