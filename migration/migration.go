package migration

import (
	"errors"
	"log"

	"gits-assignment/config"
	"gits-assignment/models"

	"gorm.io/gorm"
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

	// seed user
	if err := config.DB.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		var user models.User

		hashedPass, _ := user.HashPassword("123456")
		user.Email = "admin@gits.com"
		user.Password = hashedPass

		if err := config.DB.Create(&user).Error; err != nil {
			log.Fatalf("DATABASE::failed seeding database -> err: %v", err)
		}
	}

	// seed author
	if err := config.DB.First(&models.Author{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		var author models.Author

		author.Name = "Derek E. Malik"
		author.Email = "derekmalik@gits.com"

		if err := config.DB.Create(&author).Error; err != nil {
			log.Fatalf("DATABASE::failed seeding database -> err: %v", err)
		}
	}

	// seed book
	if err := config.DB.First(&models.Book{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		var book models.Book

		book.Title = "Atomic Habits"
		book.Year = 2023
		book.AuthorID = 1

		if err := config.DB.Create(&book).Error; err != nil {
			log.Fatalf("DATABASE::failed seeding database -> err: %v", err)
		}
	}

	// seed publisher
	if err := config.DB.First(&models.Publisher{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		var publisher models.Publisher

		publisher.Name = "Grahanusa Media"
		publisher.Location = "Jakarta"
		publisher.BookID = 1

		if err := config.DB.Create(&publisher).Error; err != nil {
			log.Fatalf("DATABASE::failed seeding database -> err: %v", err)
		}
	}
}
