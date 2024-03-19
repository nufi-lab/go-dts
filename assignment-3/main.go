package main

import (
	"assignment-3/controllers/authcontroller"
	"assignment-3/controllers/ordercontroller"
	"assignment-3/middlewares"
	"assignment-3/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	// Initialize Gin router
	r := gin.Default()

	// Define routes
	r.POST("/login", authcontroller.Login)
	r.POST("/register", authcontroller.Register)
	r.GET("/logout", authcontroller.Logout)

	// Define API route group with middleware
	api := r.Group("/api")
	api.Use(middlewares.JWTMiddleware())

	// Define routes within the API group
	api.GET("/orders", ordercontroller.Index)
	api.POST("/orders", ordercontroller.Create)
	api.PUT("/orders/:id", ordercontroller.Update)
	api.DELETE("/orders/:id", ordercontroller.Delete)

	log.Fatal(http.ListenAndServe(":8080", r))

	// r.GET("/order/:id", ordercontroller.Show)

	// r.Run(":8080")
}
