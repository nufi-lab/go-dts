package models

type Genre struct {
	GenreID uint   `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
}
