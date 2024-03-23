package models

import (
	"fmt"
	"mylib/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User struct represents the User table.
type User struct {
	gorm.Model
	ID        uint      `json:"user_id" gorm:"primaryKey"`
	RoleID    uint      `gorm:"not null"`
	FullName  string    `json:"full_name" gorm:"not null" validate:"required"`
	Username  string    `json:"username" gorm:"unique;not null" validate:"required"`
	Email     string    `json:"email" gorm:"unique;not null" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Role      Role      `json:"role" gorm:"foreignKey:RoleID"`
}

type RegisterRequest struct {
	FullName string `json:"full_name" valid:"required"`
	Username string `json:"username" valid:"required"`
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required"`
}

type LoginRequest struct {
	Username string `json:"username" valid:"required-Your username is required"`
	Password string `json:"password" valid:"required-Your password is required"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Role     string `json:"role"`
	// User        User   `json:"user"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name" valid:"required"`
	Username string `json:"username" valid:"required"`
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	if !govalidator.IsEmail(u.Email) {
		return fmt.Errorf("invalid email format")
	}

	hashedPassword, err := helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	if u.RoleID == 0 {
		u.RoleID = 2 // Assuming 1 is the default role ID
	}

	return nil
}

func (u User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	hashedPassword, err := helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	err = nil
	return
}
