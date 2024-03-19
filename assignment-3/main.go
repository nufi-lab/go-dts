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
	admin := r.Group("/admin")
	admin.Use(middlewares.JWTMiddleware(), middlewares.AuthorizeRoleMiddleware("admin"))

	// Define routes within the API group
	admin.GET("/orders", ordercontroller.Index)
	admin.POST("/orders", ordercontroller.Create)
	admin.PUT("/orders/:id", ordercontroller.Update)
	admin.DELETE("/orders/:id", ordercontroller.Delete)

	customer := r.Group("/customer")
	customer.Use(middlewares.JWTMiddleware(), middlewares.AuthorizeRoleMiddleware("customer"))

	customer.GET("/orders", ordercontroller.Index)

	log.Fatal(http.ListenAndServe(":8080", r))

	// r.GET("/order/:id", ordercontroller.Show)

	// r.Run(":8080")
}
