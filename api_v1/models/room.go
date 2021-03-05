package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID 			uint		`gorm:"primary_key" json:"id"`
	Title 		string		`gorm:"not null" json:"title"`
	Time		string		`gorm:"not null" json:"time"`
	Image 		string		`json:"img"`
	Public 		bool		`gorm:"default:false;not null" json:"public"`
}