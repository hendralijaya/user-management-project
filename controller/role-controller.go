package controller

import (
	"hendralijaya/user-management-project/exception"
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type roleController struct {
	roleService service.RoleService
	jwtService  service.JWTService
	logger helper.Log
}

func NewRoleController(roleService service.RoleService, jwtService service.JWTService, logger helper.Log) RoleController {
	return &roleController{
		roleService: roleService,
		jwtService:  jwtService,
		logger: logger,
	}
}

func (ctrl *roleController) All(context *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
	role := ctrl.roleService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   role,
	}
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := ctrl.jwtService.GetUserData(token, "user_id")
	ctrl.logger.Infof("%s already get all role data", userId)
}

func (ctrl *roleController) FindById(context *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
	notFoundError := exception.NewNotFoundError(context, ctrl.logger)
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	role, err := ctrl.roleService.FindById(uint(id))
	ok = notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   role,
	}
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	roleId, _ := ctrl.jwtService.GetUserData(token, "role_id")
	ctrl.logger.Infof("%s already get a role with id %d", roleId, id)
}

func (ctrl *roleController) Insert(context *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
	validationError := exception.NewValidationError(context, ctrl.logger)
	internalServerError := exception.NewInternalServerError(context, ctrl.logger)
	var request web.RoleCreateRequest
	err := context.BindJSON(&request)
	ok := validationError.SetMeta(err)
	if ok {
		validationError.Logf(err)
		return
	}
	role, err := ctrl.roleService.Create(request)
	ok = internalServerError.SetMeta(err)
	if ok {
		internalServerError.Logf(err)
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   role,
	}
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := ctrl.jwtService.GetUserData(token, "user_id")
	ctrl.logger.Infof("%s already created a role with id %d", userId, role.ID)
}

func (ctrl *roleController) Update(context *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
	notFoundError := exception.NewNotFoundError(context, ctrl.logger)
	validationError := exception.NewValidationError(context, ctrl.logger)
	internalServerError := exception.NewInternalServerError(context, ctrl.logger)
	var u web.RoleUpdateRequest
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := internalServerError.SetMeta(err)
	if ok {
		internalServerError.Logf(err)
		return
	}
	u.ID = id
	err = context.BindJSON(&u)
	ok = validationError.SetMeta(err)
	if ok {
		validationError.Logf(err)
		return
	}
	role, err := ctrl.roleService.Update(u)
	ok = notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   role,
	}
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := ctrl.jwtService.GetUserData(token, "user_id")
	ctrl.logger.Infof("%s already updated a role with id %d", userId, id)
}

func (ctrl *roleController) Delete(context *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
	notFoundError := exception.NewNotFoundError(context, ctrl.logger)
	internalServerError := exception.NewInternalServerError(context, ctrl.logger)
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := internalServerError.SetMeta(err)
	if ok {
		internalServerError.Logf(err)
		return
	}
	err = ctrl.roleService.Delete(uint(id))
	ok = notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   "Role has been removed",
	}
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := ctrl.jwtService.GetUserData(token, "user_id")
	ctrl.logger.Infof("%s already deleted a role with id %d", userId, id)
}
