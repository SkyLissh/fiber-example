package db

import (
	"skylissh/fiber-example/core"
	"skylissh/fiber-example/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connection() *gorm.DB {
	dsn := core.Settings.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Album{})
	db.AutoMigrate(&models.Url{})

	return db
}

var DB = connection()
