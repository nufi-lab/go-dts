package usercontroller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"mygram/models"
)

// UpdateUser handles updating user data

func UpdateUser(c *gin.Context) {
	// Parse userID from parameters
	userIDStr := c.Param("userId")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid userID"})
		return
	}

	// Parse JSON input
	var userInput models.User
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Fetch user data by userID
	var user models.User
	if err := models.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User not found"})
		return
	}

	// Update user data
	user.Email = userInput.Email
	user.Username = userInput.Username
	user.UpdatedAt = time.Now()

	// Save updated user data to database
	if err := models.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user"})
		return
	}

	// Response with updated user data
	c.JSON(http.StatusOK, gin.H{
		"id":         user.UserID,
		"email":      user.Email,
		"username":   user.Username,
		"age":        user.Age,
		"updated_at": user.UpdatedAt,
	})
}
