package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt int `gorm:"autoCreateTime" json:"created_at"`
}

func init() {
	var err error

	dsn := "default:secret@tcp(10.0.75.1:3306)/default"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}
}
