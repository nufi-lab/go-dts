package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null" valid:"required"`
	Biography string `json:"biography" gorm:"not null" valid:"required"`
}

type GetListAuthorRequest struct {
	Name      string `json:"name"`
	Biography string `json:"biography"`
}

type AuthorResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}
