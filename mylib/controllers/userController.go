package controllers

import (
	"fmt"
	"mylib/helpers"
	"mylib/middlewares"
	"mylib/models"
	"mylib/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) Routes(r *gin.RouterGroup) {
	// Fungsi helper untuk menambahkan rute ke grup
	addRoutes := func(routes func(group *gin.RouterGroup)) {
		routes(r)
	}

	// Gunakan fungsi helper untuk menambahkan rute-rute yang diinginkan
	addRoutes(func(group *gin.RouterGroup) {
		group.POST("/register", c.Register)
		group.POST("/login", c.Login)
	})

	userRouter := r.Group("/user")
	userRouter.Use(middlewares.Authentication(), middlewares.AddAuthorizationHeader())
	userRouter.PUT("/:id", middlewares.Authorization(), c.UpdateUser)
	userRouter.DELETE("/:id", middlewares.Authorization(), c.DeleteUser)

}

func (c *UserController) Register(ctx *gin.Context) {
	var request models.RegisterRequest

	err := ctx.ShouldBind(&request)

	contentType := helpers.GetContentType(ctx)
	_ = contentType

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := c.userService.Register(request)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = c.userService.LoadUserRole(&user)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user,
	})
}

func (c *UserController) Login(ctx *gin.Context) {
	var request models.LoginRequest

	err := ctx.ShouldBind(&request)

	contentType := helpers.GetContentType(ctx)
	_ = contentType

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := c.userService.Login(request)

	if err != nil {
		ctx.JSON(401, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Header("Authorization", "Bearer "+user.AccessToken)

	ctx.JSON(200, gin.H{
		"data": user,
	})
}

func (u *UserController) UpdateUser(ctx *gin.Context) {

	id := ctx.Param("id")

	convertID, err := strconv.Atoi(id)

	fmt.Println(convertID)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var request models.UpdateUserRequest

	err = ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := u.userService.UpdateUser(convertID, request)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(200, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid user ID"})
		return
	}

	err = c.userService.DeleteUser(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "User deleted successfully"})
}
