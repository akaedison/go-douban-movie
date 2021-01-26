package parse

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func getDoc(url string) *goquery.Document {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36 Edg/88.0.705.50")
	req.Header.Set("Referer", "https://accounts.douban.com/")
	req.Header.Set("Connection", "keep-alive")
	req.Cookie("ll=\"118282\"; bid=nf7xX3Arjq4; gr_user_id=9b8b3f0b-fd41-4d7e-a716-61023662ec53; _vwo_uuid_v2=DE4B60536578F882E6C518C7D166CF9E7|9647e284bf012f6adec1f90ee2e775aa; __yadk_uid=mfaFJaHtcRkWKw9eRf6i2Z7L4O8sv6At; __utmc=30149280; __utmz=30149280.1611582404.2.2.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; __utmc=223695111; __utmz=223695111.1611582404.2.2.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; ap_v=0,6.0; ucf_uid=b871236a-875a-4027-bea8-b5538b8d2a10; __gads=ID=ea126516d1f9dfbd-2290948f58c60018:T=1611582639:RT=1611582639:S=ALNI_MY1P3NgvnQUx5nbecx3EQRkcawGbA; __utma=30149280.1202205571.1609566723.1611584731.1611587513.4; __utmb=30149280.0.10.1611587513; __utmb=223695111.0.10.1611587513; __utma=223695111.1857363572.1609566729.1611584731.1611587513.4; _pk_ref.100001.4cf6=[\"\",\"\",1611587513,\"https://www.douban.com/\"]; _pk_id.100001.4cf6=f3f3ba5b419f320a.1609566729.4.1611587513.1611584731.; _pk_ses.100001.4cf6=*; GED_PLAYLIST_ACTIVITY=W3sidSI6InhQeXYiLCJ0c2wiOjE2MTE1ODc4MTgsIm52IjowLCJ1cHQiOjE2MTE1ODI0MDYsImx0IjoxNjExNTg2NjMzfV0.; dbcl2=\"164488786:l5bVUnM6C8U\"")
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

func GetDouban(url string) {
	for i := 0; i < 10; i++ {
		url := fmt.Sprintf("%v?start=%v", url, i*25)
		doc := getDoc(url)
		doc.Find(".grid_view > li > .item").Each(func(i int, selection *goquery.Selection) {
			id := selection.Find(".pic em").Eq(0).Text()
			pic, _ := selection.Find(".pic a img").Eq(0).Attr("src")
			title1 := selection.Find(".info .hd a .title").Eq(0).Text()
			title2 := selection.Find(".info .hd a .title").Eq(1).Text()
			title := fmt.Sprintf("%v%v", title1, title2)
			other := selection.Find(".info .hd a .other").Eq(0).Text()
			desc := selection.Find(".info .bd p").Eq(0).Text()
			descInfo := strings.Split(desc, "\n")
			descDirector := descInfo[0]
			movieDesc := strings.Split(descInfo[1], "/")
			//year := strings.TrimSpace(movieDesc[0])
			//area := strings.TrimSpace(movieDesc[1])
			//tag := strings.TrimSpace(movieDesc[2])
			fmt.Println(id)
			fmt.Println(pic)
			fmt.Println(title)
			fmt.Println(other)
			fmt.Println(desc)
			fmt.Println(descDirector)
			fmt.Println(movieDesc)
			//fmt.Println(year)
			//fmt.Println(area)
			//fmt.Println(tag)
		})
	}
}
