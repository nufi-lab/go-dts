package authcontroller

import (
	"fmt"
	"mylib/config"
	"mylib/helper"
	"mylib/models"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// @Summary User login
// @Description Authenticate user with provided credentials and generate access token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "User credentials"
// @Success 200 {object} models.LoginResponse "Login successful"
// @Failure 400
// @Router /login [post]
func Login(c *gin.Context) {
	var userInput models.User

	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Username or password is incorrect"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	// Check if the password is valid
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username or password is incorrect"})
		return
	}

	var roleName string

	err := config.DB.Table("roles").Where("id = ?", user.RoleID).Pluck("name", &roleName).Error
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Fetched Role Name:", roleName)

	user.Role.Name = roleName

	// Generate JWT token
	expTime := time.Now().Add(time.Minute * 10)
	claims := &config.JWTClaim{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	fmt.Println("Extracted Role:", claims.Role) // Debugging

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
	// c.Header("Authorization", "Bearer "+tokenString)
	// c.Header("Expires", expTime.Format(time.RFC1123))

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// @Summary Register a new user
// @Description Register a new user with the provided details.
// @Description If data user not found role is librarian.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.RegisterRequest true "User details"
// @Success 201 {object} models.RegisterRequest "User successfully registered"
// @Failure 400
// @Router /register [post]
func Register(c *gin.Context) {
	var userInput models.User

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

	// Hash password using bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	var existingUsersCount int64
	config.DB.Model(&models.User{}).Count(&existingUsersCount)

	// Role ID default
	var defaultRoleID uint
	if existingUsersCount > 0 {
		defaultRoleID = 2 // Jika ada pengguna di database, role ID default adalah 2
	} else {
		defaultRoleID = 1 // Jika tidak ada pengguna di database, role ID default adalah 1
	}

	userInput.RoleID = defaultRoleID

	if err := config.DB.Create(&userInput).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

// Logout godoc
// @Summary Logout
// @Description Logout the user by clearing the authentication token cookie.
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 "Logout successful"
// @Router /logout [get]
func Logout(c *gin.Context) {
	// Clear token cookie
	c.SetCookie("token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
