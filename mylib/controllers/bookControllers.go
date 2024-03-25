package controllers

import (
	"mylib/models"
	"mylib/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService *services.BookService
}

func NewBookController(BookService *services.BookService) *BookController {
	return &BookController{BookService: BookService}
}

// GetAllBooks godoc
// @Summary Get all books
// @Description Get all books
// @Tags Book
// @Accept json
// @Produce json
// @Success 200
// @Router /books [get]
func (bc *BookController) GetAllBooks(c *gin.Context) {
	var request models.GetListBookRequest
	if err := c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books, err := bc.BookService.GetAllBooks(request)
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

// GetBookById godoc
// @Summary Get book by id
// @Description Get book by id
// @Tags Book
// @Accept json
// @Produce json
// @Param id  path  string  true  "Book ID"
// @Success 200
// @Router /book/{id} [get]
func (bc *BookController) GetBookByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Book ID"})
		return
	}

	book, err := bc.BookService.FindBookByID(uint(id))
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

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the provided details
// @Tags Book
// @Accept json
// @Produce json
// @Param book body models.GetListBookRequest true "Book details"
// @Success 201 {object} models.BookResponse
// @Failure 400
// @Router /librarian/book [post]
func (bc *BookController) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the required fields
	if book.AuthorID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Author ID is required"})
		return
	}

	if book.GenreID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Genre ID is required"})
		return
	}

	if err := bc.BookService.CreateBook(&book); err != nil {
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

// UpdateBook godoc
// @Summary Update book details
// @Description Update details of an existing book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body models.GetListBookRequest true "Book details"
// @Success 200 {object} models.BookResponse
// @Failure 400
// @Router /librarian/book/{id} [put]
func (bc *BookController) UpdateBook(c *gin.Context) {
	var updateRequest models.Book
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Book ID"})
		return
	}

	book, err := bc.BookService.FindBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	book.Title = updateRequest.Title
	book.Description = updateRequest.Description
	book.PublicationYear = updateRequest.PublicationYear
	book.AvailableCopies = updateRequest.AvailableCopies

	if err := bc.BookService.UpdateBook(uint(id), book); err != nil {
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

	c.JSON(http.StatusOK, bookResponse)
}

// DeleteBook godoc
// @Summary Delete book
// @Description Delete an existing book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 204 "Book deleted successfully"
// @Failure 400 "Invalid input data"
// @Router /librarian/book/{id} [delete]
func (bc *BookController) DeleteBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Book ID"})
		return
	}

	if err := bc.BookService.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
