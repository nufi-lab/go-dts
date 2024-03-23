package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	ID   uint   `json:"genre_id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}
