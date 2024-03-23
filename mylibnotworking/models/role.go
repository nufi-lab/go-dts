package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID   uint   `json:"role_id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}
