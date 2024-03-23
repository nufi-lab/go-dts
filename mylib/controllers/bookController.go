package controllers

import (
	"assignment-3/models"
	"assignment-3/services"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService *services.BookService
}

func NewBookController(BookService *services.BookService) *BookController {
	return &BookController{BookService: BookService}
}

func (gc *BookController) GetAllBooks(c *gin.Context) {
	var request models.GetListBookRequest
	if err := c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books, err := gc.BookService.GetAllBooks(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var bookResponses []models.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, models.BookResponse{
			ID:              book.ID,
			Title:           book.Title,
			Description:     book.Description,
			Author:          book.Author.Name,
			Genre:           book.Genre.Name,
			PublicationYear: book.PublicationYear,
			AvailableCopies: book.AvailableCopies,
		})
	}

	c.JSON(http.StatusOK, bookResponses)
}

func (gc *BookController) GetBookByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Book ID"})
		return
	}

	book, err := gc.BookService.FindBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	bookResponse := models.BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		Description:     book.Description,
		Author:          book.Author.Name,
		Genre:           book.Genre.Name,
		PublicationYear: book.PublicationYear,
		AvailableCopies: book.AvailableCopies,
	}

	c.JSON(http.StatusOK, bookResponse)
}

func (gc *BookController) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi menggunakan GoValidator
	if _, err := govalidator.ValidateStruct(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := gc.BookService.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Book"})
		return
	}

	bookResponse := models.BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		Description:     book.Description,
		Author:          book.Author.Name,
		Genre:           book.Genre.Name,
		PublicationYear: book.PublicationYear,
		AvailableCopies: book.AvailableCopies,
	}

	c.JSON(http.StatusCreated, bookResponse)
}

func (gc *BookController) UpdateBook(c *gin.Context) {
	var updateRequest models.GetListBookRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Book ID"})
		return
	}

	book := &models.Book{ID: uint(id), Title: updateRequest.Title}
	if err := gc.BookService.UpdateBook(uint(id), book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Book"})
		return
	}

	bookResponse := models.BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		Description:     book.Description,
		Author:          book.Author.Name,
		Genre:           book.Genre.Name,
		PublicationYear: book.PublicationYear,
		AvailableCopies: book.AvailableCopies,
	}

	c.JSON(http.StatusCreated, bookResponse)
}

func (gc *BookController) DeleteBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Book ID"})
		return
	}

	if err := gc.BookService.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
