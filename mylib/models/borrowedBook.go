package models

import (
	"time"

	"gorm.io/gorm"
)

type BorrowedBook struct {
	gorm.Model
	ID           uint      `json:"borrowedbook_id" gorm:"primaryKey" valid:"required"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	BookID       uint      `json:"book_id" gorm:"not null" valid:"required"`
	BorrowedDate time.Time `json:"borrowed_date" gorm:"not null" valid:"required"`
	ReturnDate   time.Time `json:"return_date" gorm:"not null"`
	User         User      `json:"user" gorm:"foreignKey:UserID"`
	Book         Book      `json:"book" gorm:"foreignKey:BookID"`
}

type BorrowRequest struct {
	BookID uint `json:"book_id" binding:"required"`
	UserID uint `json:"user_id"`
}
