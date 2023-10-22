package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open("host=postgres port=5432 user=postgres password=123456 dbname=crud sslmode=disable"),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	err = db.AutoMigrate(&Person{}, &Region{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	return db
}
