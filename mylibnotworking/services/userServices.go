package services

import (
	"errors"
	"mylib/helpers"
	"mylib/models"

	"gorm.io/gorm"
)

type UserService struct {
	gorm *gorm.DB
}

func NewUserService(gorm *gorm.DB) *UserService {

	return &UserService{
		gorm: gorm,
	}
}

func (s *UserService) Register(request models.RegisterRequest) (models.User, error) {

	user := models.User{
		FullName: request.FullName,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	err := s.gorm.Create(&user).Error

	return user, err
}

func (s *UserService) LoadUserRole(user *models.User) error {
	err := s.gorm.Preload("Role").First(&user, user.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Login(request models.LoginRequest) (*models.LoginResponse, error) {
	var user models.User

	err := s.gorm.Preload("Role").Where("username = ?", request.Username).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	getPassword := user.Password

	if !helpers.ComparePassword(getPassword, request.Password) {
		return nil, errors.New("invalid username/password")
	}

	token, _ := helpers.GenerateToken(user.Username)

	return &models.LoginResponse{
		Token:    token,
		Username: user.Username,
		Role:     user.Role.Name,
	}, nil
}

func (u *UserService) GetUserByID(id int) (*models.User, error) {
	var user models.User

	err := u.gorm.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserService) UpdateUser(id int, request models.UpdateUserRequest) (*models.User, error) {
	user, err := u.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	if request.FullName != "" {
		user.FullName = request.FullName
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Username != "" {
		user.Username = request.Username
	}
	if request.Password != "" {
		hashedPassword, _ := helpers.HashPassword(request.Password)

		user.Password = hashedPassword
		// user.Password = request.Password
	}

	err = u.gorm.Save(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(userID uint) error {
	var user models.User

	err := s.gorm.First(&user, userID).Error
	if err != nil {
		return err
	}

	err = s.gorm.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}
