package database

import (
	"database/sql"
	"log"
	"mylib/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormInit(sqlDB *sql.DB) *gorm.DB {
	dsn := "user=postgres password=nfitri dbname=mylib sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	database.AutoMigrate(&models.Author{}, &models.Book{}, &models.BorrowedBook{}, &models.Genre{}, &models.Review{}, &models.User{}, &models.Role{})
	// database.AutoMigrate()

	if err != nil {
		log.Fatalf("Error when connect to database: %v", err)
	}

	return database
}
