package controller

import (
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var roleFile = "role.log"

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
		jwtService: jwtService,
	}
}

func (c *roleController) All(context *gin.Context) {
	logger := helper.NewLog(roleFile)
	role := c.roleService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   role,
	}
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := c.jwtService.GetUserId(token)
	logger.Infof("%d already get all roles", userId)
}

func (c *roleController) FindById(context *gin.Context) {
	logger := helper.NewLog(roleFile)
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	role, err := c.roleService.FindById(uint(id))
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
	userId, _ := c.jwtService.GetUserId(token)
	logger.Infof("%d already find a role data with id %d", userId, role.ID)
}

func (c *roleController) Insert(context *gin.Context) {
	logger := helper.NewLog(roleFile)
	var request web.RoleCreateRequest
	err := context.BindJSON(&request)
	ok := helper.ValidationError(context, err)
	if ok {
		return
	}
	role, err := c.roleService.Create(request)
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
	userId, _ := c.jwtService.GetUserId(token)
	logger.Infof("%d already insert a role with id %d", userId, role.ID)
}

func (c *roleController) Update(context *gin.Context) {
	
}

func (c *roleController) Delete(context *gin.Context) {
	
}
