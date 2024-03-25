package helper

import (
	"mylib/config"
	"mylib/models"
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
