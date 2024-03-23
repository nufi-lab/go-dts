package routers

import (
	"database/sql"
	"mylib/controllers"
	"mylib/middlewares"
	"mylib/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *sql.DB, gorm *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AddAuthorizationHeader())

	userService := services.NewUserService(gorm)
	userController := controllers.NewUserController(userService)
	userController.Routes(r.Group(""))

	return r

}
