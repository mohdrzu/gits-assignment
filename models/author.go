package models

import (
	"errors"
	"fmt"

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
		return nil, fmt.Errorf("failed to read author: %v", err)
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
		return fmt.Errorf("failed to create author: %v", err)
	}

	return nil
}

func (a *Author) UpdateAuthor(author Author) error {
	err := config.DB.Updates(&author).Error
	if err != nil {
		return fmt.Errorf("failed to update author: %v", err)
	}

	return nil
}

func (a *Author) DeleteAuthor(author Author) error {
	err := config.DB.Delete(&author).Error
	if err != nil {
		return fmt.Errorf("failed to delete author: %v", err)
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
