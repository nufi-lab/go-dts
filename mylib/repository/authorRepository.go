package repository

import (
	"mylib/models"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (ar *AuthorRepository) GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	if err := ar.db.Select("id, name, biography").Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func (ar *AuthorRepository) FindAuthorByID(id uint) (*models.Author, error) {
	var author models.Author
	if err := ar.db.First(&author, id).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (ar *AuthorRepository) CreateAuthor(author *models.Author) error {
	return ar.db.Create(author).Error
}

func (ar *AuthorRepository) UpdateAuthor(id uint, author *models.Author) error {
	return ar.db.Model(&models.Author{}).Where("id = ?", id).Updates(&author).Error
}

func (ar *AuthorRepository) DeleteAuthor(id uint) error {
	return ar.db.Delete(&models.Author{}, id).Error
}
