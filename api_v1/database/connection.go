package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

var (
	// DBConn is a pointer to gorm.DB
	DBConn		*gorm.DB
	user 		= os.Getenv("MYSQL_USER")
	password	= os.Getenv("MYSQL_PASSWORD")
	host		= os.Getenv("MYSQL_HOST")
	dbname 		= os.Getenv("MYSQL_DATABASE")
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