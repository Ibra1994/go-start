package migrations

import (
	"github.com/Ibra1994/go-start/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDatabase() {
	dsn := "host=10.0.75.1 user=default password=secret dbname=default port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.Organization{})

	organization := models.Organization{Name: "Root"}
	db.Create(&organization)
}
