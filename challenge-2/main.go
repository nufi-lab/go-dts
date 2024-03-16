package main

import (
	"challenge-2/controllers/ordercontroller"
	"challenge-2/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.POST("/orders", ordercontroller.Create)
	r.GET("/orders", ordercontroller.Index)
	// r.GET("/order/:id", ordercontroller.Show)
	r.PUT("/orders/:id", ordercontroller.Update)
	r.DELETE("/orders/:id", ordercontroller.Delete)

	r.Run(":8080")
}
