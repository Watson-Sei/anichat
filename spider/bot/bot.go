package bot

import (
	"fmt"
	"github.com/Watson-Sei/anichat/spider/plugin"
	"log"
	"net/http"
	"time"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID     uint   `gorm:"primary_key" json:"id"`
	Title  string `gorm:"not null" json:"title"`
	Time   string `gorm:"not null" json:"time"`
	Image  string `json:"img"`
	Public bool   `gorm:"default:false; not null" json:"public"`
}

func Scraping(db *gorm.DB) {

	dbInitialize(db)

	//db.Where("Public = ?", false).Delete(&Room{})
	//db.Where("Public = ?", true).Delete(&Room{})
	//log.Println("DBテーブルを初期化しました")

	day := time.Now()
	const layout = "20060102"

	url := fmt.Sprintf("https://anime.eiga.com/tv/q/?b=2&d=%s&p=13", day.Format(layout))
	//url := "https://anime.eiga.com/tv/q/?b=2&d=20210403&p=13"

	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	// NewDocumentFromReader
	// https://pkg.go.dev/github.com/PuerkitoBio/goquery#NewDocumentFromReader
	doc, _ := goquery.NewDocumentFromReader(res.Body)

	// アニメ・声優　番組表（テレビ番組) FInd The Title
	fmt.Println(doc.Find(".handTtlL").Text())

	// 最新に情報を取得します。
	doc.Find(".dayProgram").Each(func(i int, s *goquery.Selection) {
		s.Find("li").Each(func(l1 int, s1 *goquery.Selection) {
			if s1.Find(".tvScheTime").Text() != "" {
				// アニメタイトルがあるか
				if s1.Find(".tvScheTtl").Text() != "" {
					// アニメイメージ
					s1.Find("img").Each(func(_ int, s *goquery.Selection) {
						url, _ := s.Attr("src")
						if strings.Contains(url, "jpg") {
							fmt.Printf("Image: %s \n", url)

							// DBに追加する
							room := Room{Title: s1.Find(".tvScheTtl").Text(), Time: s1.Find(".tvScheTime").Text(), Image: url}
							db.Create(&room)

							// AT Command Set
							err := plugin.Reformat(s1.Find(".tvScheTime").Text(), room.ID)
							if err != nil {
								panic(err)
							}
						}
					})
				}
			}
		})
	})
}

func dbInitialize(db *gorm.DB)  {
	db.Where("Public = ?", true).Delete(&Room{})
	db.Where("Public = ?", false).Delete(&Room{})
	log.Printf("番組DBを初期化しました。")
}