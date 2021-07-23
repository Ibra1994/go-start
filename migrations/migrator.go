package migrations

import (
	"github.com/Ibra1994/go-start/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDatabase() {
	dsn := "default:secret@tcp(10.0.75.1:3306)/default"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.Organization{})

	organization := models.Organization{Name: "Root"}
	db.Create(&organization)
}
