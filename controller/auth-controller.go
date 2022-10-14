package controller

import (
	"hendralijaya/user-management-project/exception"
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/service"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var userManagementLog = "user-management.log"

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	ForgotPasswordEmail(ctx *gin.Context)
	ForgotPassword(ctx *gin.Context)
	VerifyRegisterToken(ctx *gin.Context)
	VerifyForgotPasswordToken(ctx *gin.Context)
}

type authController struct {
	userService service.UserService
	jwtService  service.JWTService
	authService service.AuthService
	logger 	   helper.Log
}

func NewAuthController(authService service.AuthService, userService service.UserService, jwtService service.JWTService, logger helper.Log) AuthController {
	return &authController{
		authService: authService,
		userService: userService,
		jwtService:  jwtService,
		logger: logger,
	}
}

func (ctrl *authController) Login(ctx *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
	authError := exception.NewAuthenticationError(ctx, ctrl.logger)
	validationError := exception.NewValidationError(ctx, ctrl.logger)
	internalServerError := exception.NewInternalServerError(ctx, ctrl.logger)
	var u web.UserLoginRequest
	err := ctx.BindJSON(&u)
	ok := validationError.SetMeta(err)
	if ok {
		ctrl.logger.Errorf("Login failed with error: %v", err)
		return
	}
	user, err := ctrl.authService.Login(u)
	ok = authError.SetMeta(err)
	if ok {
		ctrl.logger.Infof("%d failed to login", user.ID)
		return
	}
	generateToken, err := ctrl.jwtService.GenerateToken(strconv.FormatUint(uint64(user.ID), 10), user.Username, user.Email, uint(user.RoleId), 60)
	ok = internalServerError.SetMeta(err)
	if ok {
		internalServerError.Logf(err)
		return
	}
	if user.VerificationTime.IsZero() {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Errors: "Please verify your account",
			Data:   nil,
		}
		ctx.JSON(http.StatusOK, webResponse)
		return
	}
	user.AuthToken = generateToken
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data:   user,
	}
	ctx.JSON(http.StatusOK, webResponse)
	ctrl.logger.Infof("%d already login", user.ID)
}

func (ctrl *authController) Register(ctx *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
	validationError := exception.NewValidationError(ctx, ctrl.logger)
	internalServerError := exception.NewInternalServerError(ctx, ctrl.logger)
	var u web.UserRegisterRequest
	err := ctx.BindJSON(&u)
	ok := validationError.SetMeta(err)
	if ok {
		validationError.Logf(err)
		return
	}
	u.RoleId = 2
	u.VerificationTime = time.Now()
	user, err := ctrl.authService.Register(u)
	ok = internalServerError.SetMeta(err)
	if ok {
		ctrl.logger.Errorf("Register failed with error: %v", err)
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Success",
		Errors: nil,
		Data:   user,
	}
	ctx.JSON(http.StatusCreated, webResponse)
	ctrl.logger.Infof("%d already registered", user.ID)
}

func (ctrl *authController) ForgotPasswordEmail(ctx *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
	validationError := exception.NewValidationError(ctx, ctrl.logger)
	internalServerError := exception.NewInternalServerError(ctx, ctrl.logger)
	var u web.UserForgotPasswordRequest
	err := ctx.BindJSON(&u)
	ok := validationError.SetMeta(err)
	if ok {
		validationError.Logf(err)
		return
	}
	user := ctrl.userService.FindByEmail(u.Email)
	token, err := ctrl.jwtService.GenerateToken(strconv.FormatUint(uint64(user.ID), 10), user.Username, user.Email, uint(user.RoleId), 60*24)
	ok = internalServerError.SetMeta(err)
	if ok {
		internalServerError.Logf(err)
		return
	}
	mainLink := helper.GetMainLink()
	helper.SendMail(`<a href="`+mainLink+`/auth/verify_forgot_password_token/`+token+`">Click this link</a>`, "Forgot Password Email", user.Email, user.Email, user.FirstName+" "+user.LastName)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data:   user,
	}
	ctx.JSON(http.StatusOK, webResponse)
	ctrl.logger.Infof("%d already send the forgot password email", user.ID)
}

func (ctrl *authController) ForgotPassword(ctx *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
	validationError := exception.NewValidationError(ctx, ctrl.logger)
	internalServerError := exception.NewInternalServerError(ctx, ctrl.logger)
	var u web.UserForgotPasswordRequest
	err := ctx.BindJSON(&u)
	ok := validationError.SetMeta(err)
	if ok {
		validationError.Logf(err)
		return
	}
	user := ctrl.userService.FindByUsername(u.Username)
	token, err := ctrl.jwtService.GenerateToken(strconv.FormatUint(uint64(user.ID), 10), user.Username, user.Email, uint(user.RoleId), 60*24)
	ok = internalServerError.SetMeta(err)
	if ok {
		internalServerError.Logf(err)
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data:   token,
	}
	ctx.JSON(http.StatusOK, webResponse)
	ctrl.logger.Infof("%d already send the forgot password email", user.ID)
}

func (ctrl *authController) VerifyRegisterToken(ctx *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
	internalServerError := exception.NewInternalServerError(ctx, ctrl.logger)
	notFoundError := exception.NewNotFoundError(ctx, ctrl.logger)
	tokenError := exception.NewTokenError(ctx, ctrl.logger)
	userToken := ctx.Param("token")
	jwtToken, err := ctrl.jwtService.ValidateToken(userToken)
	ok := tokenError.SetMeta(err)
	if ok {
		tokenError.Logf(err)
		return
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	userIdString := claims["user_id"].(string)
	userId, err := strconv.ParseUint(userIdString, 10, 64)
	ok = internalServerError.SetMeta(err)
	if ok {
		internalServerError.Logf(err)
		return
	}
	user, err := ctrl.userService.FindById(uint(userId))
	ok = notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	var userRequest web.UserRegisterVerificationTokenRequest
	userRequest.VerificationTime = time.Now()
	userRequest.ID = userId
	user, err = ctrl.authService.VerifyRegisterToken(userRequest)
	ok = internalServerError.SetMeta(err)
	if ok {
		internalServerError.Logf(err)
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data:   user,
	}
	ctx.JSON(http.StatusOK, webResponse)
	ctrl.logger.Infof("%d already verify registered token", user.ID)
}

func (ctrl *authController) VerifyForgotPasswordToken(ctx *gin.Context) {
	ctrl.logger.SetUp(userManagementLog)
	internalServerError := exception.NewInternalServerError(ctx, ctrl.logger)
	notFoundError := exception.NewNotFoundError(ctx, ctrl.logger)
	tokenError := exception.NewTokenError(ctx, ctrl.logger)
	validationError := exception.NewValidationError(ctx, ctrl.logger)
	var u web.UserNewPasswordRequest
	err := ctx.BindJSON(&u)
	ok := validationError.SetMeta(err)
	if ok {
		validationError.Logf(err)
		return
	}
	userToken := ctx.Param("token")
	jwtToken, err := ctrl.jwtService.ValidateToken(userToken)
	ok = tokenError.SetMeta(err)
	if ok {
		tokenError.Logf(err)
		return
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	userIdString := claims["user_id"].(string)
	userId, err := strconv.ParseUint(userIdString, 10, 64)
	ok = internalServerError.SetMeta(err)
	if ok {
		internalServerError.Logf(err)
		return
	}
	user, err := ctrl.userService.FindById(uint(userId))
	ok = notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	u.ID = userId
	user, err = ctrl.authService.VerifyForgotPasswordToken(u)
	ok = notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data:   user,
	}
	ctx.JSON(http.StatusOK, webResponse)
	ctrl.logger.Infof("%d already verify forgot password token", user.ID)
}
