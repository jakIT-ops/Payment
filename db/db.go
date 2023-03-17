package db

import (
	"gorm.io/driver/postgres"
	"log"
	"payment_full/models"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dbURL := "postgres://jakit:pass@localhost:5433/payment"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	//db, err := sql.Open("postgres", "user=myuser password=mypassword dbname=mydatabase sslmode=disable")
	if err != nil {
		log.Panic("Failed to connect database")
	}
	db.AutoMigrate(&models.Account{}, &models.Transaction{}, &models.User{})

	Database = DbInstance{
		Db: db,
	}
}
