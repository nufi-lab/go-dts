package helper

import (
	"assignment-3/config"
	"assignment-3/models"
)

func IsEmailExists(email string) (bool, error) {
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		// Handle error
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func IsUsernameExists(username string) (bool, error) {
	var user models.User
	result := config.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		// Handle error
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
