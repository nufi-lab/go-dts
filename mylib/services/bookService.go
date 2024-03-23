package services

import (
	"assignment-3/models"
	"assignment-3/repository"
)

type BookService struct {
	bookRepo *repository.BookRepository
}

func NewBookService(bookRepo *repository.BookRepository) *BookService {
	return &BookService{bookRepo: bookRepo}
}

func (gs *BookService) GetAllBooks(request models.GetListBookRequest) ([]models.Book, error) {
	return gs.bookRepo.GetAllBooks()
}

func (gs *BookService) FindBookByID(id uint) (*models.Book, error) {
	return gs.bookRepo.FindBookByID(id)
}

func (gs *BookService) CreateBook(Book *models.Book) error {
	return gs.bookRepo.CreateBook(Book)
}

func (gs *BookService) UpdateBook(id uint, Book *models.Book) error {
	return gs.bookRepo.UpdateBook(id, Book)
}

func (gs *BookService) DeleteBook(id uint) error {
	return gs.bookRepo.DeleteBook(id)
}
