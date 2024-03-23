package controllers

import (
	"fmt"
	"mylib/helpers"
	"mylib/middlewares"
	"mylib/models"
	"mylib/services"
	"net/http"
	"strconv"
	"time"

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

// func (c *UserController) Routes(r *gin.RouterGroup) {
// 	addRoutes := func(routes func(group *gin.RouterGroup)) {
// 		routes(r)
// 	}

// 	addRoutes(func(group *gin.RouterGroup) {
// 		group.POST("/register", c.Register)
// 		group.POST("/login", c.Login)
// 	})

// 	userRouter := r.Group("/user")
// 	userRouter.Use(middlewares.Authenticate, middlewares.AddAuthorizationHeader())
// 	userRouter.PUT("/:id", middlewares.Authorization(), c.UpdateUser)
// 	userRouter.DELETE("/:id", middlewares.Authorization(), c.DeleteUser)

// }

func (c *UserController) Routes(r *gin.RouterGroup) {
	// Routes for registration and login
	r.POST("/register", c.Register)
	r.POST("/login", c.Login)

	// Routes for authenticated actions
	r.PUT("/user/:id", middlewares.Authenticate(), middlewares.AddAuthorizationHeader(), c.UpdateUser)
	r.DELETE("/user/:id", middlewares.Authenticate(), middlewares.AddAuthorizationHeader(), c.DeleteUser)
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

	ctx.JSON(http.StatusCreated, gin.H{
		"user_id":   user.ID,
		"email":     user.Email,
		"username":  user.Username,
		"full_name": user.FullName,
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
	expTime := time.Now().Add(time.Hour)

	// Set header HTTP untuk menyimpan token
	ctx.Header("Authorization", "Bearer "+user.Token)
	ctx.Header("Expires", expTime.Format(time.RFC1123))

	// ctx.Header("Authorization", "Bearer "+user.Token)

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
