package config

import (
	"petlover/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=containers-us-west-40.railway.app user=postgres  dbname=railway port=7212 password=9tSADWCzD3on9SLPWbTW sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	migration()
}

func migration() {
	// automate
	DB.AutoMigrate(&models.Pets{})
}
