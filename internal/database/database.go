package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabse() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection")

	dbUser := os.Getenv("DB_USERNAME")
	dbPwd := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbTable := os.Getenv("DB_TABLE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPwd, dbTable, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return db, err
	}

	if err := sqlDb.Ping(); err != nil {
		return db, err
	}

	return db, nil
}
