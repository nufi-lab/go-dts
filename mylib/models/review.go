package models

import "time"

type Review struct {
	ReviewID  uint      `gorm:"primaryKey"`
	BookID    uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	Rating    uint      `gorm:"not null"`
	Comment   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	User      User      `gorm:"foreignKey:UserID"`
	Book      Book      `gorm:"foreignKey:BookID"`
}
