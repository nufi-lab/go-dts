package models

import (
	"time"

	"gorm.io/gorm"
)

type BorrowedBook struct {
	gorm.Model
	ID           uint      `json:"borrowedbook_id" gorm:"primaryKey"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	BookID       uint      `json:"book_id" gorm:"not null"`
	BorrowedDate time.Time `json:"borrowed_date" gorm:"not null"`
	ReturnDate   time.Time `json:"return_date" gorm:"not null"`
	User         User      `json:"user" gorm:"foreignKey:UserID"`
	Book         Book      `json:"book" gorm:"foreignKey:BookID"`
}
