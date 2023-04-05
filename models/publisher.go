package models

import (
	"errors"

	"gits-assignment/config"

	"gorm.io/gorm"
)

type Publisher struct {
	gorm.Model
	Name     string
	Location string
	BookID   uint
}

func (p *Publisher) ReadPublishers() ([]Publisher, error) {
	var pub []Publisher
	err := config.DB.Find(&pub).Error
	if err != nil {
		return nil, errors.New("failed to read publishers")
	}

	return pub, nil
}

func (p *Publisher) CreatePublisher(name, location string, book uint) error {
	newPub := Publisher{
		Name:     name,
		Location: location,
		BookID:   book,
	}

	err := config.DB.Create(&newPub).Error
	if err != nil {
		return errors.New("failed to create publisher")
	}

	return nil
}

func (p *Publisher) UpdatePublisher(pub Publisher) error {
	err := config.DB.Updates(&pub).Error
	if err != nil {
		return errors.New("failed to update publisher")
	}

	return nil
}

func (p *Publisher) DeletePublisher(pub Publisher) error {
	err := config.DB.Delete(&pub).Error
	if err != nil {
		return errors.New("failed to delete publisher")
	}

	return nil
}

func (p *Publisher) FindPublisherByID(id int) (Publisher, error) {
	var pub Publisher
	result := config.DB.Where("id", id).Find(&pub)
	if result.RowsAffected < 1 {
		return Publisher{}, errors.New("record not found")
	}

	return pub, nil
}
