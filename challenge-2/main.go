package main

import (
	"challenge-2/controllers/ordercontroller"
	"challenge-2/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/orders", ordercontroller.Index)
	r.GET("/order/:id", ordercontroller.Show)
	r.POST("/order", ordercontroller.Create)
	r.PUT("/order/:id", ordercontroller.Update)
	r.DELETE("/order", ordercontroller.Delete)

	r.Run()
}
