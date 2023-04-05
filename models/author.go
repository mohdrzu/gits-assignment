package models

import (
	"errors"

	"gits-assignment/config"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name  string
	Email string
	Books []Book
}

func (a *Author) ReadAuthors() ([]Author, error) {
	var author []Author
	err := config.DB.Model(&Author{}).Preload("Books").Find(&author).Error
	if err != nil {
		return nil, errors.New("failed to read authors")
	}

	return author, nil
}

func (a *Author) CreateAuthor(name, email string) error {
	newAuthor := Author{
		Name:  name,
		Email: email,
	}

	err := config.DB.Create(&newAuthor).Error
	if err != nil {
		return errors.New("failed to create authors")
	}

	return nil
}

func (a *Author) UpdateAuthor(author Author) error {
	err := config.DB.Updates(&author).Error
	if err != nil {
		return errors.New("failed to update authors")
	}

	return nil
}

func (a *Author) DeleteAuthor(author Author) error {
	err := config.DB.Delete(&author).Error
	if err != nil {
		return errors.New("failed to delete authors")
	}

	return nil
}

func (a *Author) FindAuthorByID(id int) (Author, error) {
	var author Author
	result := config.DB.Where("id", id).Find(&author)
	if result.RowsAffected < 1 {
		return Author{}, errors.New("record not found")
	}

	return author, nil
}
