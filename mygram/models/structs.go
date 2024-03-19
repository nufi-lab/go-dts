package models

import "time"

// User struct represents the User table.
type User struct {
	ID        uint
	Username  string
	Email     string
	Password  string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Photo struct represents the Photo table.
type Photo struct {
	ID        uint // Primary key
	Title     string
	Caption   string
	PhotoURL  string
	UserID    uint // Foreign key of User Table
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Comment struct represents the Comment table.
type Comment struct {
	ID        uint // Primary key
	UserID    uint // Foreign key of User Table
	PhotoID   uint // Foreign key of Photo Table
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// SocialMedia struct represents the SocialMedia table.
type SocialMedia struct {
	ID             uint // Primary key
	Name           string
	SocialMediaURL string
	UserID         uint // Foreign key of User Table
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
