package controller

import (
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var userFile = "user-management.log"

type UserController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (ctrl *userController) All(context *gin.Context) {
	logger := helper.NewLog(userFile)
	users := ctrl.userService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   users,
	}
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := ctrl.jwtService.GetUserId(token)
	logger.Infof("%d already get all users", userId)
}

func (ctrl *userController) FindById(context *gin.Context) {
	logger := helper.NewLog(userFile)
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	user, err := ctrl.userService.FindById(uint(id))
	ok = helper.NotFoundError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   user,
	}
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := ctrl.jwtService.GetUserId(token)
	logger.Infof("%d already find a user data with id %d", userId, user.ID)
}

func (ctrl *userController) Insert(context *gin.Context) {
	logger := helper.NewLog(userFile)
	var u web.UserCreateRequest
	err := context.BindJSON(&u)
	ok := helper.ValidationError(context, err)
	if ok {
		return
	}
	u.VerificationTime = time.Now()
	user, err := ctrl.userService.Create(u)
	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   user,
	}
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := ctrl.jwtService.GetUserId(token)
	logger.Infof("%d already insert a user with id %d", userId, user.ID)
}

func (ctrl *userController) Update(context *gin.Context) {
	logger := helper.NewLog(userFile)
	var u web.UserUpdateRequest
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	u.ID = id
	err = context.BindJSON(&u)
	ok = helper.ValidationError(context, err)
	if ok {
		return
	}
	user, err := ctrl.userService.Update(u)
	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   user,
	}
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := ctrl.jwtService.GetUserId(token)
	logger.Infof("%d already updated a user with id %d", userId, user.ID)
}

func (ctrl *userController) Delete(context *gin.Context) {
	logger := helper.NewLog(userFile)
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	err = ctrl.userService.Delete(uint(id))
	ok = helper.NotFoundError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   "User has been removed",
	}
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := ctrl.jwtService.GetUserId(token)
	logger.Infof("%d already deleted a user with id %d", userId, id)
}
