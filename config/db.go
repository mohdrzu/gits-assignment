package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func setDSN() string {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbName, port)

	return dsn
}

func InitDatabase() {
	dsn := setDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatalf("DATABASE::failed to connect to database -> err: %v", err)
	}

	DB = db
}
