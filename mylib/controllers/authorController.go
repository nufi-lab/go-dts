package controllers

import (
	"assignment-3/models"
	"assignment-3/services"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	authorService *services.AuthorService
}

func NewAuthorController(authorService *services.AuthorService) *AuthorController {
	return &AuthorController{authorService: authorService}
}

func (gc *AuthorController) GetAllAuthors(c *gin.Context) {
	var request models.GetListAuthorRequest
	if err := c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authors, err := gc.authorService.GetAllAuthors(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var authorResponses []models.AuthorResponse
	for _, Author := range authors {
		authorResponses = append(authorResponses, models.AuthorResponse{
			ID:        Author.ID,
			Name:      Author.Name,
			Biography: Author.Biography,
		})
	}

	c.JSON(http.StatusOK, authorResponses)
}

func (gc *AuthorController) GetAuthorByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Author ID"})
		return
	}

	author, err := gc.authorService.FindAuthorByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	authorResponse := models.AuthorResponse{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
	}

	c.JSON(http.StatusOK, authorResponse)
}

func (gc *AuthorController) CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi menggunakan GoValidator
	if _, err := govalidator.ValidateStruct(author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := gc.authorService.CreateAuthor(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create author"})
		return
	}

	authorResponse := models.AuthorResponse{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
	}

	c.JSON(http.StatusCreated, authorResponse)
}

func (gc *AuthorController) UpdateAuthor(c *gin.Context) {
	var updateRequest models.GetListAuthorRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Author ID"})
		return
	}

	author := &models.Author{ID: uint(id), Name: updateRequest.Name}
	if err := gc.authorService.UpdateAuthor(uint(id), author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Author"})
		return
	}
	authorResponse := models.AuthorResponse{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
	}

	c.JSON(http.StatusCreated, authorResponse)
}

func (gc *AuthorController) DeleteAuthor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Author ID"})
		return
	}

	if err := gc.authorService.DeleteAuthor(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}
