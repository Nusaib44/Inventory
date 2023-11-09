package db

import (
	models "inventory/pkg/models/Database"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(url string) {
	println("database  url in db init", url)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	DB = db

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Return{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Provider{})
	db.AutoMigrate(&models.Payment{})
	db.AutoMigrate(&models.Area{})
}
