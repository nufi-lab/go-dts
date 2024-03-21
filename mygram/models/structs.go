package models

import "time"

// User struct represents the User table.
type User struct {
	UserID    uint      `json:"user_id" gorm:"primaryKey;column:user_id"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password"`
	Age       int       `json:"age" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Photo struct represents the Photo table.
type Photo struct {
	PhotoID   uint      `json:"photo_id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url" gorm:"not null"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Comment struct represents the Comment table.
type Comment struct {
	CommentID uint      `json:"comment_id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	PhotoID   uint      `json:"photo_id"`
	Photo     Photo     `json:"photo" gorm:"foreignKey:PhotoID"`
	Message   string    `json:"message" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SocialMedia struct represents the SocialMedia table.
type SocialMedia struct {
	SocialMediaID  uint      `json:"social_media_id" gorm:"primaryKey"`
	Name           string    `json:"name" gorm:"not null"`
	SocialMediaURL string    `json:"social_media_url" gorm:"not null"`
	UserID         uint      `json:"user_id"`
	User           User      `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
