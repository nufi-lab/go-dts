package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null"`
	Biography string `json:"biography" gorm:"not null"`
}
