package repository

import (
	"assignment-3/models"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (gr *AuthorRepository) GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	if err := gr.db.Select("id, name").Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func (gr *AuthorRepository) FindAuthorByID(id uint) (*models.Author, error) {
	var author models.Author
	if err := gr.db.First(&author, id).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (gr *AuthorRepository) CreateAuthor(author *models.Author) error {
	return gr.db.Create(author).Error
}

func (gr *AuthorRepository) UpdateAuthor(id uint, author *models.Author) error {
	return gr.db.Model(&models.Author{}).Where("id = ?", id).Updates(&author).Error
}

func (gr *AuthorRepository) DeleteAuthor(id uint) error {
	return gr.db.Delete(&models.Author{}, id).Error
}
