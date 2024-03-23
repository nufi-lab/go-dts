package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID              uint   `json:"book_id" gorm:"primaryKey"`
	Title           string `json:"tittle" gorm:"not null"`
	AuthorID        uint   `json:"author_id" gorm:"not null"`
	GenreID         uint   `json:"genre_id" gorm:"not null"`
	Description     string `json:"description" gorm:"not null"`
	PublicationYear uint   `json:"publication_year" gorm:"not null"`
	AvailableCopies uint   `json:"available_copies" gorm:"not null"`
	Author          Author `json:"author" gorm:"foreignKey:AuthorID"`
	Genre           Genre  `json:"genre" gorm:"foreignKey:GenreID"`
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
