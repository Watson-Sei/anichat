package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"time"
)

func main()  {
	day := time.Now()
	const layout = "20060102"

	url := fmt.Sprintf("https://anime.eiga.com/tv/q/?b=2&d=%s&p=13", day.Format(layout))

	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	// NewDocumentFromReader
	// https://pkg.go.dev/github.com/PuerkitoBio/goquery#NewDocumentFromReader
	doc, _ := goquery.NewDocumentFromReader(res.Body)

	// アニメ・声優 番組表（テレビ番組）Find The Title
	fmt.Println(doc.Find(".headTtlL").Text())

	// 最新日付の情報を取得します。
	doc.Find(".dayProgram").Each(func(i int, s *goquery.Selection) {
		//
		s.Find("li").Each(func(i1 int, s1 *goquery.Selection) {
			if s1.Find(".tvScheTime").Text() != "" {
				// アニメのタイトル（テキスト）
				fmt.Printf("Title: %s \n ", s1.Find(".tvScheTtl").Text())
				// アニメの開始と終了時間(テキスト)
				fmt.Printf("Time: %s \n", s1.Find(".tvScheTime").Text())
				// アニメのイメージリンク（テキスト）
				s1.Find("img").Each(func(_ int, s *goquery.Selection) {
					url, _ := s.Attr("src")
					if strings.Contains(url, "jpg") {
						fmt.Printf("Img: %s \n", url)
					}
				})
			}
		})
	})
}