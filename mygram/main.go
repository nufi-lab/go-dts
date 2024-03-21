package main

import (
	"log"
	"mygram/controllers/authcontroller"
	"mygram/controllers/usercontroller"
	"mygram/middlewares"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()

	r.POST("/users/login", authcontroller.Login)
	r.POST("/users/register", authcontroller.Register)

	r.PUT("/users/:userId", middlewares.JWTMiddleware(), usercontroller.UpdateUser)

	// r.GET("/logout", authcontroller.Logout)
	log.Fatal(http.ListenAndServe(":8080", r))

}
