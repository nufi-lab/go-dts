package services

import (
	"mylib/models"
	"mylib/repository"
)

type BookService struct {
	bookRepo *repository.BookRepository
}

func NewBookService(bookRepo *repository.BookRepository) *BookService {
	return &BookService{bookRepo: bookRepo}
}

func (bs *BookService) GetAllBooks(request models.GetListBookRequest) ([]models.Book, error) {
	return bs.bookRepo.GetAllBooks()
}

func (bs *BookService) FindBookByID(id uint) (*models.Book, error) {
	return bs.bookRepo.FindBookByID(id)
}

func (bs *BookService) CreateBook(Book *models.Book) error {
	return bs.bookRepo.CreateBook(Book)
}

func (bs *BookService) UpdateBook(id uint, Book *models.Book) error {
	return bs.bookRepo.UpdateBook(id, Book)
}

func (gs *BookService) DeleteBook(id uint) error {
	return gs.bookRepo.DeleteBook(id)
}
