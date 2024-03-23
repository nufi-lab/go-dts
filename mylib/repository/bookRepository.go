package repository

import (
	"assignment-3/models"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (gr *BookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	if err := gr.db.Select("id, name").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (gr *BookRepository) FindBookByID(id uint) (*models.Book, error) {
	var book models.Book
	if err := gr.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (gr *BookRepository) CreateBook(book *models.Book) error {
	return gr.db.Create(book).Error
}

func (gr *BookRepository) UpdateBook(id uint, book *models.Book) error {
	return gr.db.Model(&models.Book{}).Where("id = ?", id).Updates(&book).Error
}

func (gr *BookRepository) DeleteBook(id uint) error {
	return gr.db.Delete(&models.Book{}, id).Error
}
