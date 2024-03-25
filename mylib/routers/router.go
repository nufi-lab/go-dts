package routers

import (
	"mylib/controllers"
	"mylib/controllers/authcontroller"
	_ "mylib/docs"
	"mylib/middlewares"
	"mylib/repository"
	"mylib/services"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @title           Swagger My Lib
// @version         1.0
// @description     Documentation for mylib.

// @host      localhost:8080
// @BasePath  /
func StartServer(gorm *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/login", authcontroller.Login)
	r.POST("/register", authcontroller.Register)
	r.GET("/logout", authcontroller.Logout)
	publicRouter := r.Group("")

	librarian := r.Group("/librarian")
	librarian.Use(middlewares.JWTMiddleware(), middlewares.AuthorizeRoleMiddleware("librarian"))

	genreRepo := repository.NewGenreRepository(gorm)
	genreService := services.NewGenreService(genreRepo)
	genreController := controllers.NewGenreController(genreService)

	publicRouter.GET("/genres", genreController.GetAllGenres)
	publicRouter.GET("/genre/:id", genreController.GetGenreByID)
	librarian.POST("/genre", genreController.CreateGenre)
	librarian.PUT("/genre/:id", genreController.UpdateGenre)
	librarian.DELETE("/genre/:id", genreController.DeleteGenre)

	authorRepo := repository.NewAuthorRepository(gorm)
	authorService := services.NewAuthorService(authorRepo)
	authorController := controllers.NewAuthorController(authorService)

	publicRouter.GET("/authors", authorController.GetAllAuthors)
	publicRouter.GET("/author/:id", authorController.GetAuthorByID)
	librarian.POST("/author", authorController.CreateAuthor)
	librarian.PUT("/author/:id", authorController.UpdateAuthor)
	librarian.DELETE("/author/:id", authorController.DeleteAuthor)

	bookRepo := repository.NewBookRepository(gorm)
	bookService := services.NewBookService(bookRepo)
	bookController := controllers.NewBookController(bookService)

	publicRouter.GET("/books", bookController.GetAllBooks)
	publicRouter.GET("/book/:id", bookController.GetBookByID)
	librarian.POST("/book", bookController.CreateBook)
	librarian.PUT("/book/:id", bookController.UpdateBook)
	librarian.DELETE("/book/:id", bookController.DeleteBook)

	visitor := r.Group("/visitor")
	visitor.Use(middlewares.JWTMiddleware(), middlewares.AuthorizeRoleMiddleware("visitor"))

	visitor.POST("/borrow-book", controllers.BorrowBook)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
