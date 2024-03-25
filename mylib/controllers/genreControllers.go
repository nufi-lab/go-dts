package controllers

import (
	"mylib/models"
	"mylib/services"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type GenreController struct {
	genreService *services.GenreService
}

func NewGenreController(genreService *services.GenreService) *GenreController {
	return &GenreController{genreService: genreService}
}

// GetAllGenres godoc
// @Summary get all genres
// @Description get all genres
// @Tags Genre
// @Accept json
// @Produce json
// @Success 200
// @Router /genres [get]
func (gc *GenreController) GetAllGenres(c *gin.Context) {
	var request models.GetListGenreRequest
	if err := c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	genres, err := gc.genreService.GetAllGenres(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var genreResponses []models.GenreResponse
	for _, genre := range genres {
		genreResponses = append(genreResponses, models.GenreResponse{
			ID:   genre.ID,
			Name: genre.Name,
		})
	}

	c.JSON(http.StatusOK, genreResponses)
}

// GetGenreById godoc
// @Summary get genre by id
// @Description get genre by id
// @Tags Genre
// @Accept json
// @Produce json
// @Param id  path  string  true  "Genre ID"
// @Success 200
// @Router /genre/{id} [get]
func (gc *GenreController) GetGenreByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
		return
	}

	genre, err := gc.genreService.FindGenreByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
		return
	}

	genreResponse := models.GenreResponse{
		ID:   genre.ID,
		Name: genre.Name,
	}

	c.JSON(http.StatusOK, genreResponse)
}

// CreateGenre godoc
// @Summary Create a new genre
// @Description Create a new genre with the provided details
// @Tags Genre
// @Accept json
// @Produce json
// @Param genre body models.GetListGenreRequest true "Genre details"
// @Success 201 {object} models.GenreResponse
// @Failure 400
// @Router /librarian/genre [post]
func (gc *GenreController) CreateGenre(c *gin.Context) {
	var genre models.Genre
	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi menggunakan GoValidator
	if _, err := govalidator.ValidateStruct(genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := gc.genreService.CreateGenre(&genre); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create genre"})
		return
	}

	genreResponse := models.GenreResponse{
		ID:   genre.ID,
		Name: genre.Name,
	}

	c.JSON(http.StatusCreated, genreResponse)
}

// UpdateGenre godoc
// @Summary Update genre details
// @Description Update details of an existing genre
// @Tags Genre
// @Accept json
// @Produce json
// @Param id path string true "Genre ID"
// @Param genre body models.GetListGenreRequest true "Genre details"
// @Success 200 {object} models.GenreResponse
// @Failure 400
// @Router /librarian/genre/{id} [put]
func (gc *GenreController) UpdateGenre(c *gin.Context) {
	var updateRequest models.GetListGenreRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
		return
	}

	genre := &models.Genre{ID: uint(id), Name: updateRequest.Name}
	if err := gc.genreService.UpdateGenre(uint(id), genre); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update genre"})
		return
	}
	genreResponse := models.GenreResponse{
		ID:   genre.ID,
		Name: genre.Name,
	}

	c.JSON(http.StatusCreated, genreResponse)
}

// DeleteGenre godoc
// @Summary Delete genre
// @Description Delete an existing genre
// @Tags Genre
// @Accept json
// @Produce json
// @Param id path string true "Genre ID"
// @Success 204 "Genre deleted successfully"
// @Failure 400 "Invalid input data"
// @Router /librarian/genre/{id} [delete]
func (gc *GenreController) DeleteGenre(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
		return
	}

	if err := gc.genreService.DeleteGenre(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete genre"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Genre deleted successfully"})
}

// func Index(c *gin.Context) {
// 	var genres []models.Genre
// 	config.DB.Preload("Genres").Find(&genres)

// 	var filteredGenres []models.Genre
// 	for _, genre := range genres {
// 		if genre.ID != 0 {
// 			filteredGenres = append(filteredGenres, genre)
// 		}
// 	}

// 	c.JSON(http.StatusOK, filteredGenres)
// }

// func Create(c *gin.Context) {
// 	var requestData struct {
// 		Name string `json:"name"`
// 	}

// 	if err := c.BindJSON(&requestData); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}

// 	// Validate the Name field using govalidator
// 	if valid := govalidator.IsNotNull(requestData.Name); !valid {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Name field cannot be empty"})
// 		return
// 	}

// 	genre := models.Genre{
// 		Name: requestData.Name,
// 	}

// 	config.DB.Create(&genre)

// 	c.JSON(http.StatusCreated, gin.H{"message": "Genre created successfully", "genre": gin.H{
// 		"id":   genre.ID,
// 		"name": genre.Name,
// 	}})
// }

// func Update(c *gin.Context) {
// 	genreID := c.Param("id")

// 	var updatedGenreData struct {
// 		Name string `json:"name"`
// 	}
// 	if err := c.BindJSON(&updatedGenreData); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var existingGenre models.Genre
// 	if err := config.DB.First(&existingGenre, genreID).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
// 		return
// 	}

// 	if updatedGenreData.Name != "" {
// 		existingGenre.Name = updatedGenreData.Name
// 	}

// 	if err := config.DB.Save(&existingGenre).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update genre"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Genre updated successfully", "genre": gin.H{
// 		"id":   existingGenre.ID,
// 		"name": existingGenre.Name,
// 	}})
// }

// func Delete(c *gin.Context) {
// 	genreID := c.Param("id")

// 	var existingGenre models.Genre
// 	if err := config.DB.First(&existingGenre, genreID).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
// 		return
// 	}

// 	if err := config.DB.Delete(&existingGenre).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Genre deleted successfully"})
// }
