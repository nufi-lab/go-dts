package services

import (
	"mylib/models"
	"mylib/repository"
)

type AuthorService struct {
	authorRepo *repository.AuthorRepository
}

func NewAuthorService(authorRepo *repository.AuthorRepository) *AuthorService {
	return &AuthorService{authorRepo: authorRepo}
}

func (as *AuthorService) GetAllAuthors(request models.GetListAuthorRequest) ([]models.Author, error) {
	return as.authorRepo.GetAllAuthors()
}

func (as *AuthorService) FindAuthorByID(id uint) (*models.Author, error) {
	return as.authorRepo.FindAuthorByID(id)
}

func (as *AuthorService) CreateAuthor(author *models.Author) error {
	return as.authorRepo.CreateAuthor(author)
}

func (as *AuthorService) UpdateAuthor(id uint, author *models.Author) error {
	return as.authorRepo.UpdateAuthor(id, author)
}

func (as *AuthorService) DeleteAuthor(id uint) error {
	return as.authorRepo.DeleteAuthor(id)
}
