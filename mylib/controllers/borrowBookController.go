package controllers

import (
	"mylib/config"
	"mylib/middlewares"
	"mylib/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// BorrowBook godoc
// @Summary Borrow a book
// @Description Borrow a book by providing book ID. The user must be authenticated.
// @Tags Book
// @Accept json
// @Produce json
// @Param borrowRequest body models.BorrowRequest true "Borrow request details"
// @Success 201 "Book borrowed successfully"
// @Failure 400
// @Router /borrow-book [post]
func BorrowBook(c *gin.Context) {
	userID, err := middlewares.VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var borrowRequest models.BorrowRequest
	if err := c.BindJSON(&borrowRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if borrowRequest.BookID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book ID is required"})
		return
	}

	var book models.Book
	if err := config.DB.First(&book, borrowRequest.BookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if book.AvailableCopies == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book is not available for borrowing"})
		return
	}

	borrowedDate := time.Now()
	returnDate := borrowedDate.AddDate(0, 0, 7)

	book.AvailableCopies--
	if err := config.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book availability"})
		return
	}

	borrowedBook := models.BorrowedBook{
		BookID:       book.ID,
		UserID:       userID, // Menggunakan userID dari token JWT
		BorrowedDate: borrowedDate,
		ReturnDate:   returnDate,
	}
	if err := config.DB.Create(&borrowedBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create borrowed book record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book borrowed successfully", "book_id": book.ID, "borrowed_date": borrowedDate, "return_date": returnDate})
}
