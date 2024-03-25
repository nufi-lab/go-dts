package repository

import (
	"mylib/models"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (br *BookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	if err := br.db.Preload("Author").Preload("Genre").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (br *BookRepository) FindBookByID(id uint) (*models.Book, error) {
	var book models.Book
	if err := br.db.Preload("Author").Preload("Genre").First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (br *BookRepository) CreateBook(book *models.Book) error {
	if err := br.db.Preload("Author").Preload("Genre").FirstOrCreate(book, book).Error; err != nil {
		return err
	}
	return nil
}

func (br *BookRepository) UpdateBook(id uint, book *models.Book) error {
	updateData := map[string]interface{}{
		"title":            book.Title,
		"description":      book.Description,
		"publication_year": book.PublicationYear,
		"available_copies": book.AvailableCopies,
		"author_id":        book.AuthorID,
		"genre_id":         book.GenreID,
	}

	if err := br.db.Model(&models.Book{}).Where("id = ?", id).Updates(updateData).Error; err != nil {
		// Jika terjadi kesalahan, coba lakukan pembaruan tanpa memperbarui kolom-kolom yang berhubungan dengan kunci asing
		return br.db.Model(&models.Book{}).Where("id = ?", id).Updates(book).Error
	}
	return nil
}

func (br *BookRepository) DeleteBook(id uint) error {
	return br.db.Delete(&models.Book{}, id).Error
}
