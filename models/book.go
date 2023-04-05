package models

import (
	"errors"
	"fmt"

	"gits-assignment/config"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title     string
	Year      uint
	AuthorID  uint
	Publisher *Publisher `json:",omitempty"`
}

func (b *Book) ReadBooks() ([]Book, error) {
	var book []Book
	err := config.DB.Model(&Book{}).Preload("Publisher").Find(&book).Error
	if err != nil {
		return nil, fmt.Errorf("failed to read book: %v", err)
	}

	return book, nil
}

func (b *Book) CreateBook(title string, year, author uint) error {
	newBook := Book{
		Title:    title,
		Year:     year,
		AuthorID: author,
	}

	err := config.DB.Create(&newBook).Error
	if err != nil {
		return fmt.Errorf("failed to create book: %v", err)
	}

	return nil
}

func (b *Book) UpdateBook(book Book) error {
	err := config.DB.Updates(&book).Error
	if err != nil {
		return fmt.Errorf("failed to update book: %v", err)
	}

	return nil
}

func (b *Book) DeleteBook(book Book) error {
	err := config.DB.Delete(&book).Error
	if err != nil {
		return fmt.Errorf("failed to delete book: %v", err)
	}

	return nil
}

func (b *Book) FindBookByID(id int) (Book, error) {
	var book Book
	result := config.DB.Where("id", id).Find(&book)
	if result.RowsAffected < 1 {
		return Book{}, errors.New("record not found")
	}

	return book, nil
}
