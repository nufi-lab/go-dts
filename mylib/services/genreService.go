package services

import (
	"mylib/models"
	"mylib/repository"
)

type GenreService struct {
	genreRepo *repository.GenreRepository
}

func NewGenreService(genreRepo *repository.GenreRepository) *GenreService {
	return &GenreService{genreRepo: genreRepo}
}

func (gs *GenreService) GetAllGenres(request models.GetListGenreRequest) ([]models.Genre, error) {
	return gs.genreRepo.GetAllGenres()
}

func (gs *GenreService) FindGenreByID(id uint) (*models.Genre, error) {
	return gs.genreRepo.FindGenreByID(id)
}

func (gs *GenreService) CreateGenre(genre *models.Genre) error {
	return gs.genreRepo.CreateGenre(genre)
}

func (gs *GenreService) UpdateGenre(id uint, genre *models.Genre) error {
	return gs.genreRepo.UpdateGenre(id, genre)
}

func (gs *GenreService) DeleteGenre(id uint) error {
	return gs.genreRepo.DeleteGenre(id)
}
