package models

import "time"

type BorrowedBook struct {
	BorrowedBookID uint      `gorm:"primaryKey"`
	UserID         uint      `gorm:"not null"`
	BookID         uint      `gorm:"not null"`
	BorrowedDate   time.Time `gorm:"not null"`
	ReturnDate     time.Time `gorm:"not null"`
	User           User      `gorm:"foreignKey:UserID"`
	Book           Book      `gorm:"foreignKey:BookID"`
}
