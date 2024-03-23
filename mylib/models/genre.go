package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null" valid:"required"`
}

type GetListGenreRequest struct {
	Name string `json:"name"`
}

type GenreResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
