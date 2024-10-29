package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("../database/database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("database connected")

	return db, nil
}
