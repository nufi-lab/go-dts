package repository

import (
	"mylib/models"

	"gorm.io/gorm"
)

type GenreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) *GenreRepository {
	return &GenreRepository{db: db}
}

func (gr *GenreRepository) GetAllGenres() ([]models.Genre, error) {
	var genres []models.Genre
	if err := gr.db.Select("id, name").Find(&genres).Error; err != nil {
		return nil, err
	}
	return genres, nil
}

func (gr *GenreRepository) FindGenreByID(id uint) (*models.Genre, error) {
	var genre models.Genre
	if err := gr.db.First(&genre, id).Error; err != nil {
		return nil, err
	}
	return &genre, nil
}

func (gr *GenreRepository) CreateGenre(genre *models.Genre) error {
	return gr.db.Create(genre).Error
}

func (gr *GenreRepository) UpdateGenre(id uint, genre *models.Genre) error {
	return gr.db.Model(&models.Genre{}).Where("id = ?", id).Updates(&genre).Error
}

func (gr *GenreRepository) DeleteGenre(id uint) error {
	return gr.db.Delete(&models.Genre{}, id).Error
}
