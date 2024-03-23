package routers

import (
	"assignment-3/controllers"
	"assignment-3/controllers/authcontroller"
	"assignment-3/middlewares"
	"assignment-3/repository"
	"assignment-3/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(gorm *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/login", authcontroller.Login)
	r.POST("/register", authcontroller.Register)
	r.GET("/logout", authcontroller.Logout)

	librarian := r.Group("/librarian")
	librarian.Use(middlewares.JWTMiddleware(), middlewares.AuthorizeRoleMiddleware("librarian"))

	genreRepo := repository.NewGenreRepository(gorm)
	genreService := services.NewGenreService(genreRepo)
	genreController := controllers.NewGenreController(genreService)

	librarian.GET("/genres", genreController.GetAllGenres)
	librarian.GET("/genre/:id", genreController.GetGenreByID)
	librarian.POST("/genre", genreController.CreateGenre)
	librarian.PUT("/genre/:id", genreController.UpdateGenre)
	librarian.DELETE("/genre/:id", genreController.DeleteGenre)

	authorRepo := repository.NewAuthorRepository(gorm)
	authorService := services.NewAuthorService(authorRepo)
	authorController := controllers.NewAuthorController(authorService)

	librarian.GET("/authors", authorController.GetAllAuthors)
	librarian.GET("/author/:id", authorController.GetAuthorByID)
	librarian.POST("/author", authorController.CreateAuthor)
	librarian.PUT("/author/:id", authorController.UpdateAuthor)
	librarian.DELETE("/author/:id", authorController.DeleteAuthor)

	return r
}
