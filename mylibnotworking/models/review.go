package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ID      uint   `json:"review_id" gorm:"primaryKey"`
	BookID  uint   `json:"book_id" gorm:"not null"`
	UserID  uint   `json:"user_id" gorm:"not null"`
	Rating  uint   `json:"rating" gorm:"not null"`
	Comment string `json:"comment" gorm:"not null"`
	User    User   `json:"user" gorm:"foreignKey:UserID"`
	Book    Book   `json:"book" gorm:"foreignKey:BookID"`
}
