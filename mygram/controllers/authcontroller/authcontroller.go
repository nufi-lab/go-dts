package authcontroller

import (
	"mygram/config"
	helper "mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var userInput models.User

	// Parse JSON input
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Validate email format
	if !govalidator.IsEmail(userInput.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email format"})
		return
	}

	// Validate email uniqueness
	exists, _ := helper.IsEmailExists(userInput.Email)
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
		return
	}

	// Validate non-empty and unique username
	if govalidator.IsNull(userInput.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username cannot be empty"})
		return
	}

	exists, _ = helper.IsUsernameExists(userInput.Username)
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username already exists"})
		return
	}

	// Validate non-empty and length of password
	if govalidator.IsNull(userInput.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password cannot be empty"})
		return
	}

	if !govalidator.MinStringLength(userInput.Password, strconv.Itoa(6)) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password must be at least 6 characters long"})
		return
	}

	// Validate non-empty and minimum age
	if userInput.Age < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Age must be at least 8"})
		return
	}

	// Hash password using bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// Insert into database
	if err := models.DB.Create(&userInput).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  201,
		"message": "Registration successful",
		"data": gin.H{
			"age":      userInput.Age,
			"email":    userInput.Email,
			"id":       userInput.UserID,
			"username": userInput.Username,
		},
	})

}

func Login(c *gin.Context) {
	var userInput models.User

	// Parse JSON input
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if govalidator.IsNull(userInput.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email cannot be empty"})
		return
	}

	if govalidator.IsNull(userInput.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password cannot be empty"})
		return
	}

	// Fetch user data by username
	var user models.User
	if err := models.DB.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	// Check if the password is valid
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Password incorrect"})
		return
	}

	// Generate JWT token
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JWT_KEY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Set token in cookie
	// Calculate the duration until the expiration time
	expirationDuration := time.Until(expTime)

	// Convert the duration to seconds
	expirationSeconds := int(expirationDuration.Seconds())

	// Set the cookie with expiration time
	c.SetCookie("token", tokenString, expirationSeconds, "/", "", false, true)

	// Response with token
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Login successful",
		"data": gin.H{
			"token": tokenString,
		},
	})
}

func Logout(c *gin.Context) {
	// Clear token cookie
	c.SetCookie("token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
