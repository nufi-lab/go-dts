package models

type Author struct {
	AuthorID  uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Biography string `gorm:"not null"`
}
