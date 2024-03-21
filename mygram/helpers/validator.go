package helper

import (
	"mygram/models"
)

func IsEmailExists(email string) (bool, error) {
	var user models.User
	result := models.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		// Handle error
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func IsUsernameExists(username string) (bool, error) {
	var user models.User
	result := models.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		// Handle error
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
