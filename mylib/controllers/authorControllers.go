package controllers

import (
	"mylib/models"
	"mylib/services"
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

// GetAllAuthors godoc
// @Summary Get all authors
// @Description Get all authors
// @Tags Author
// @Accept json
// @Produce json
// @Success 200
// @Router /authors [get]
func (ac *AuthorController) GetAllAuthors(c *gin.Context) {
	var request models.GetListAuthorRequest
	if err := c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authors, err := ac.authorService.GetAllAuthors(request)
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

// GetAuthorById godoc
// @Summary Get author by id
// @Description Get author by id
// @Tags Author
// @Accept json
// @Produce json
// @Param id  path  string  true  "Author ID"
// @Success 200
// @Router /author/{id} [get]
func (ac *AuthorController) GetAuthorByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Author ID"})
		return
	}

	author, err := ac.authorService.FindAuthorByID(uint(id))
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

// CreateAuthor godoc
// @Summary Create a new author
// @Description Create a new author with the provided details
// @Tags Author
// @Accept json
// @Produce json
// @Param author body models.GetListAuthorRequest true "Author details"
// @Success 201 {object} models.AuthorResponse
// @Failure 400
// @Router /librarian/author [post]
func (ac *AuthorController) CreateAuthor(c *gin.Context) {
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

	if err := ac.authorService.CreateAuthor(&author); err != nil {
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

// UpdateAuthor godoc
// @Summary Update author details
// @Description Update details of an existing author
// @Tags Author
// @Accept json
// @Produce json
// @Param id path string true "Author ID"
// @Param author body models.GetListAuthorRequest true "Author details"
// @Success 200 {object} models.AuthorResponse
// @Failure 400
// @Router /librarian/author/{id} [put]
func (ac *AuthorController) UpdateAuthor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Author ID"})
		return
	}

	var updateRequest models.GetListAuthorRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author, err := ac.authorService.FindAuthorByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Author"})
		return
	}

	if updateRequest.Name != "" {
		author.Name = updateRequest.Name
	}
	if updateRequest.Biography != "" {
		author.Biography = updateRequest.Biography
	}

	if err := ac.authorService.UpdateAuthor(uint(id), author); err != nil {
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

// DeleteAuthor godoc
// @Summary Delete author
// @Description Delete an existing author
// @Tags Author
// @Accept json
// @Produce json
// @Param id path string true "Author ID"
// @Success 204 "Author deleted successfully"
// @Failure 400 "Invalid input data"
// @Router /librarian/author/{id} [delete]
func (ac *AuthorController) DeleteAuthor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Author ID"})
		return
	}

	if err := ac.authorService.DeleteAuthor(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}
