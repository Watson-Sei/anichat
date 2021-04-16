package  main

import (
	"flag"
	"fmt"
	"github.com/Watson-Sei/anichat/spider/bot"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn   *gorm.DB
	user     = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	host     = os.Getenv("MYSQL_HOST")
	dbname   = os.Getenv("MYSQL_DATABASE")
)

func Connect() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbname)
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := DBConn.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}

func main() {
	var (
		mode    = flag.Int("mode", 0, "mode")
		process = flag.Int("p", 0, "process mode")
		id      = flag.Int("id", 0, "primary_key id")
	)
	flag.Parse()
	switch *mode {
	case 0:
		fmt.Println("デフォルトです")
	case 1:
		fmt.Println("スクレイピンです")
		if err := Connect(); err != nil {
			log.Panic("Can't connect database: ", err.Error())
		}
		bot.Scraping(DBConn)
		fmt.Println("スクレイピングが終了しました。")
	case 2:
		fmt.Println("cronコントローラーです")
		if err := Connect(); err != nil {
			log.Panic("Can't connect database: ", err.Error())
		}
		var room bot.Room
		switch *process {
		case 0:
			fmt.Println("処理プロセスを定義してください。")
		case 1:
			fmt.Println("Start Chat Cron!!")
			DBConn.Where("ID = ?", *id).First(&room)
			room.Public = true
			DBConn.Save(&room)
		case 2:
			fmt.Println("Finish Chat Cron!!")
			DBConn.Where("ID = ?", *id).First(&room)
			room.Public = false
			DBConn.Save(&room)
		}
	case 3:
		fmt.Println("削除コマンドテスト")
		if err := Connect(); err != nil {
			log.Panic("Can't connect database: ", err.Error())
		}
		DBConn.Where("Public = ?", false).Delete(&bot.Room{})
		DBConn.Where("Public = ?", true).Delete(&bot.Room{})
		log.Println("DBテーブルを初期化しました")
	}

	fmt.Println("コマンド完了")
}