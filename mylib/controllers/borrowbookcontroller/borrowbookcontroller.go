package borrowbookcontroller

import (
	"assignment-3/config"
	"assignment-3/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func BorrowBook(c *gin.Context) {
	var borrowRequest struct {
		BookID uint `json:"book_id" binding:"required"`
		UserID uint `json:"user_id" binding:"required"`
	}

	if err := c.BindJSON(&borrowRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if borrowRequest.BookID == 0 || borrowRequest.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book ID and User ID are required"})
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
		UserID:       borrowRequest.UserID,
		BorrowedDate: borrowedDate,
		ReturnDate:   returnDate,
	}
	if err := config.DB.Create(&borrowedBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create borrowed book record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book borrowed successfully", "book_id": book.ID, "borrowed_date": borrowedDate, "return_date": returnDate})
}
