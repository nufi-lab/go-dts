package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID              uint   `json:"book_id" gorm:"primaryKey"`
	Title           string `json:"title" gorm:"not null" valid:"required"`
	AuthorID        uint   `json:"author_id" gorm:"not null" valid:"required"`
	GenreID         uint   `json:"genre_id" gorm:"not null" valid:"required"`
	Description     string `json:"description" gorm:"not null" valid:"required"`
	PublicationYear uint   `json:"publication_year" gorm:"not null" valid:"required"`
	AvailableCopies uint   `json:"available_copies" gorm:"not null" valid:"required"`
	Author          Author `json:"author" gorm:"foreignKey:AuthorID"`
	Genre           Genre  `json:"genre" gorm:"foreignKey:GenreID"`
}

type GetListBookRequest struct {
	Title           string `json:"title"`
	AuthorID        uint   `json:"author_id"`
	GenreID         uint   `json:"genre_id"`
	Description     string `json:"description"`
	PublicationYear uint   `json:"publication_year"`
	AvailableCopies uint   `json:"available_copies"`
}

type BookResponse struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Genre           string `json:"genre"`
	Description     string `json:"description"`
	PublicationYear uint   `json:"publication_year"`
	AvailableCopies uint   `json:"available_copies"`
}
