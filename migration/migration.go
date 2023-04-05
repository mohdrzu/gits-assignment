package migration

import (
	"log"

	"gits-assignment/config"
	"gits-assignment/models"
)

func SyncDatabase() {
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Author{},
		&models.Book{},
		&models.Publisher{},
	)
	if err != nil {
		log.Fatalf("DATABASE::failed migrating database -> err: %v", err)
	}
}

func SeedDatabase() {

}
