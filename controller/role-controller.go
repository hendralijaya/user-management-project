package controller

import (
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var roleFile = "user-management.log"

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
}

func NewRoleController(roleService service.RoleService, jwtService service.JWTService) RoleController {
	return &roleController{
		roleService: roleService,
		jwtService:  jwtService,
	}
}

func (ctrl *roleController) All(context *gin.Context) {
	logger := helper.NewLog(roleFile)
	role := ctrl.roleService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   role,
	}
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := ctrl.jwtService.GetUserId(token)
	logger.Infof("%d already get all roles", userId)
}

func (ctrl *roleController) FindById(context *gin.Context) {
	logger := helper.NewLog(roleFile)
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	role, err := ctrl.roleService.FindById(uint(id))
	ok = helper.NotFoundError(context, err)
	if ok {
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
	roleId, _ := ctrl.jwtService.GetUserId(token)
	logger.Infof("%d already find a role data with id %d", roleId, role.ID)
}

func (ctrl *roleController) Insert(context *gin.Context) {
	logger := helper.NewLog(roleFile)
	var request web.RoleCreateRequest
	err := context.BindJSON(&request)
	ok := helper.ValidationError(context, err)
	if ok {
		return
	}
	role, err := ctrl.roleService.Create(request)
	ok = helper.InternalServerError(context, err)
	if ok {
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
	userId, _ := ctrl.jwtService.GetUserId(token)
	logger.Infof("%d already insert a role with id %d", userId, role.ID)
}

func (ctrl *roleController) Update(context *gin.Context) {
	logger := helper.NewLog(roleFile)
	var u web.RoleUpdateRequest
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
	role, err := ctrl.roleService.Update(u)
	ok = helper.InternalServerError(context, err)
	if ok {
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
	roleId, _ := ctrl.jwtService.GetUserId(token)
	logger.Infof("%d already updated a role with id %d", roleId, role.ID)
}

func (ctrl *roleController) Delete(context *gin.Context) {
	logger := helper.NewLog(roleFile)
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	err = ctrl.roleService.Delete(uint(id))
	ok = helper.NotFoundError(context, err)
	if ok {
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
	roleId, _ := ctrl.jwtService.GetUserId(token)
	logger.Infof("%d already deleted a role with id %d", roleId, id)
}
