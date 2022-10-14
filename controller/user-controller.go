package controller

import (
	"hendralijaya/user-management-project/exception"
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/service"
	"net/http"
	"strconv"
	"time"

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
	jwtService  service.JWTService
	logger 	helper.Log
}

func NewUserController(userService service.UserService, jwtService service.JWTService, logger helper.Log) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
		logger: logger,
	}
}

func (ctrl *userController) All(context *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
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
	ctrl.logger.Infof("%d already get all users", userId)
}

func (ctrl *userController) FindById(context *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
	notFoundError := exception.NewNotFoundError(context, ctrl.logger)
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	user, err := ctrl.userService.FindById(uint(id))
	ok = notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
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
	ctrl.logger.Infof("%d already find a user data with id %d", userId, user.ID)
}

func (ctrl *userController) Insert(context *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
	validationError := exception.NewValidationError(context, ctrl.logger)
	internalServerError := exception.NewInternalServerError(context, ctrl.logger)
	var u web.UserCreateRequest
	err := context.BindJSON(&u)
	ok := validationError.SetMeta(err)
	if ok {
		validationError.Logf(err)
		return
	}
	u.VerificationTime = time.Now()
	user, err := ctrl.userService.Create(u)
	ok = internalServerError.SetMeta(err)
	if ok {
		internalServerError.Logf(err)
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
	ctrl.logger.Infof("%d already insert a user with id %d", userId, user.ID)
}

func (ctrl *userController) Update(context *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
	validatioError := exception.NewValidationError(context, ctrl.logger)
	notFoundError := exception.NewNotFoundError(context, ctrl.logger)
	var u web.UserUpdateRequest
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	u.ID = id
	err = context.BindJSON(&u)
	ok = validatioError.SetMeta(err)
	if ok {
		validatioError.Logf(err)
		return
	}
	user, err := ctrl.userService.Update(u)
	ok = notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
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
	ctrl.logger.Infof("%d already updated a user with id %d", userId, user.ID)
}

func (ctrl *userController) Delete(context *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
	notFoundError := exception.NewNotFoundError(context, ctrl.logger)
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	err = ctrl.userService.Delete(uint(id))
	ok = notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
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
	ctrl.logger.Infof("%d already deleted a user with id %d", userId, id)
}
