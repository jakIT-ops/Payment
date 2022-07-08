package db

import (
	"log"
	"payment_full/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dbURL := "postgres://pg:pass@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect database")
	}

	db.AutoMigrate(&models.Account{}, &models.Transaction{}, &models.User{})

	Database = DbInstance{
		Db: db,
	}
}
