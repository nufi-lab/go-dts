package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Book struct {
	BookID          uint      `gorm:"primaryKey"`
	Title           string    `gorm:"not null"`
	AuthorID        uint      `gorm:"not null"`
	GenreID         uint      `gorm:"not null"`
	Description     string    `gorm:"not null"`
	PublicationYear uint      `gorm:"not null"`
	AvailableCopies uint      `gorm:"not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	Author          Author    `gorm:"foreignKey:AuthorID"`
	Genre           Genre     `gorm:"foreignKey:GenreID"`
}

func (b Book) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(b)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (b Book) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(b)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
