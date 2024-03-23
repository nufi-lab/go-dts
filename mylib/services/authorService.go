package services

import (
	"assignment-3/models"
	"assignment-3/repository"
)

type AuthorService struct {
	authorRepo *repository.AuthorRepository
}

func NewAuthorService(authorRepo *repository.AuthorRepository) *AuthorService {
	return &AuthorService{authorRepo: authorRepo}
}

func (gs *AuthorService) GetAllAuthors(request models.GetListAuthorRequest) ([]models.Author, error) {
	return gs.authorRepo.GetAllAuthors()
}

func (gs *AuthorService) FindAuthorByID(id uint) (*models.Author, error) {
	return gs.authorRepo.FindAuthorByID(id)
}

func (gs *AuthorService) CreateAuthor(Author *models.Author) error {
	return gs.authorRepo.CreateAuthor(Author)
}

func (gs *AuthorService) UpdateAuthor(id uint, Author *models.Author) error {
	return gs.authorRepo.UpdateAuthor(id, Author)
}

func (gs *AuthorService) DeleteAuthor(id uint) error {
	return gs.authorRepo.DeleteAuthor(id)
}
