package controller

import (
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
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

func (c *userController) All(context *gin.Context) {
	users := c.userService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Errors: "",
		Data:   users,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (c *userController) FindByID(context *gin.Context) {
	idString := context.Param("id")
	id , err := strconv.ParseUint(idString, 10, 64)
	user, err := c.userService.FindById(id)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Errors: "",
		Data:   user,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (c *userController) Insert(context *gin.Context) {
	var u web.UserRegisterRequest
	err := context.BindJSON(&u)
	ok := helper.ValidationError(context, err)
	if ok {
		return
	}
	u.Role_id = 1
	user, err := c.userService.Create(u)
	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Errors: "",
		Data:   user,
	}
	context.JSON(http.StatusOK, webResponse)
}