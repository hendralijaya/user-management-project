package controller

import (
	"hendralijaya/user-management-project/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService: jwtService,
	}
}
