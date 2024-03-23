package main

import (
	"assignment-3/config"
	"assignment-3/routers"
)

func main() {

	// Connect to the database
	config.ConnectDatabase()

	// Get the database instance from the config package
	db := config.DB

	start := routers.StartServer(db)
	start.Run(":8080")

	// // Initialize Gin router
	// r := gin.Default()

	// // Define routes
	// r.POST("/login", authcontroller.Login)
	// r.POST("/register", authcontroller.Register)
	// r.GET("/logout", authcontroller.Logout)

	// librarian := r.Group("/librarian")
	// librarian.Use(middlewares.JWTMiddleware(), middlewares.AuthorizeRoleMiddleware("librarian"))

	// librarian.GET("/genre", genrecontroller.Index)
	// librarian.POST("/genre", genrecontroller.Create)
	// librarian.PUT("/genre/:id", genrecontroller.Update)
	// librarian.DELETE("/genre/:id", genrecontroller.Delete)

	// librarian.GET("/author", authorcontroller.Index)
	// librarian.POST("/author", authorcontroller.Create)
	// librarian.PUT("/author/:id", authorcontroller.Update)
	// librarian.DELETE("/author/:id", authorcontroller.Delete)

	// librarian.GET("/book", bookcontroller.Index)
	// librarian.POST("/book", bookcontroller.Create)
	// librarian.PUT("/book/:id", bookcontroller.Update)
	// librarian.DELETE("/book/:id", bookcontroller.Delete)

	// visitor := r.Group("/visitor")
	// visitor.Use(middlewares.JWTMiddleware(), middlewares.AuthorizeRoleMiddleware("visitor"))

	// visitor.GET("/genre", genrecontroller.Index)

	// log.Fatal(http.ListenAndServe(":8080", r))

}
