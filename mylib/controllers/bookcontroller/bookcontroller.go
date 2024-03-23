package bookcontroller

import (
	"assignment-3/config"
	"assignment-3/models"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

var requestBookData struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	PublicationYear uint   `json:"publication_year"`
	AvailableCopies uint   `json:"available_copies"`
	AuthorID        uint   `json:"author_id"`
	GenreID         uint   `json:"genre_id"`
}

func Index(c *gin.Context) {
	var books []models.Book
	config.DB.Preload("Author").Preload("Genre").Find(&books)

	var responseBooks []gin.H
	for _, book := range books {
		responseBook := gin.H{
			"book_id":          book.ID,
			"title":            book.Title,
			"author_name":      book.Author.Name, // Assuming Name is the field representing author's name
			"genre_name":       book.Genre.Name,  // Assuming Name is the field representing genre's name
			"description":      book.Description,
			"publication_year": book.PublicationYear,
			"available_copies": book.AvailableCopies,
		}

		responseBooks = append(responseBooks, responseBook)
	}

	c.JSON(http.StatusOK, responseBooks)
}

func Create(c *gin.Context) {

	if err := c.BindJSON(&requestBookData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if valid := govalidator.IsNotNull(requestBookData.Title); !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title field cannot be empty"})
		return
	}

	if valid := govalidator.IsNotNull(requestBookData.Description); !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Description field cannot be empty"})
		return
	}

	publicationYearStr := strconv.FormatUint(uint64(requestBookData.PublicationYear), 10)
	if valid := govalidator.IsNotNull(publicationYearStr); !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PublicationYear field cannot be empty"})
		return
	}

	if requestBookData.AvailableCopies == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Available Copies field cannot be zero"})
		return
	}

	authorIDStr := strconv.FormatUint(uint64(requestBookData.AuthorID), 10)
	if valid := govalidator.IsNotNull(authorIDStr); !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Author field cannot be empty"})
		return
	}

	genreIDStr := strconv.FormatUint(uint64(requestBookData.GenreID), 10)
	if valid := govalidator.IsNotNull(genreIDStr); !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Genre field cannot be empty"})
		return
	}

	Book := models.Book{
		Title:           requestBookData.Title,
		Description:     requestBookData.Description,
		PublicationYear: requestBookData.PublicationYear,
		AvailableCopies: requestBookData.AvailableCopies,
		AuthorID:        requestBookData.AuthorID,
		GenreID:         requestBookData.GenreID,
	}

	config.DB.Create(&Book)

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "Book": gin.H{
		"id":               Book.ID,
		"name":             Book.Title,
		"description":      Book.Description,
		"publication_year": Book.PublicationYear,
		"available_copies": Book.AvailableCopies,
	}})
}

func Update(c *gin.Context) {
	bookID := c.Param("id")

	if err := c.BindJSON(&requestBookData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingBook models.Book
	if err := config.DB.First(&existingBook, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if requestBookData.Title == existingBook.Title && requestBookData.Description == existingBook.Description {
		c.JSON(http.StatusOK, gin.H{"message": "Not updated, data remains the same"})
		return
	}

	if requestBookData.Title != "" {
		existingBook.Title = requestBookData.Title
	}

	if requestBookData.Description != "" {
		existingBook.Description = requestBookData.Description
	}

	if requestBookData.PublicationYear != 0 {
		existingBook.PublicationYear = requestBookData.PublicationYear
	}

	if requestBookData.AvailableCopies != 0 {
		existingBook.AvailableCopies = requestBookData.AvailableCopies
	}

	if err := config.DB.Save(&existingBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully", "Book": gin.H{
		"id":          existingBook.ID,
		"title":       existingBook.Title,
		"description": existingBook.Description,
	}})
}

func Delete(c *gin.Context) {
	bookID := c.Param("id")

	var existingBook models.Book
	if err := config.DB.First(&existingBook, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := config.DB.Delete(&existingBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
