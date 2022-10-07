package controller

import (
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/service"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var authFile = "user-management.log"

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
}

func NewAuthController(authService service.AuthService, userService service.UserService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		userService: userService,
		jwtService:  jwtService,
	}
}

func (ctrl *authController) Login(ctx *gin.Context) {
	var u web.UserLoginRequest
	err := ctx.BindJSON(&u)
	ok := helper.ValidationError(ctx, err)
	if ok {
		return
	}
	user, err := ctrl.authService.Login(u)
	ok = helper.AuthenticationError(ctx, err)
	if ok {
		return
	}
	generateToken, err := ctrl.jwtService.GenerateToken(strconv.FormatUint(uint64(user.ID), 10), user.Username, user.Email, uint(user.RoleId), 60)
	ok = helper.InternalServerError(ctx, err)
	if ok {
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
	logger := helper.NewLog(authFile)
	logger.Infof("%d already login", user.ID)
}

func (ctrl *authController) Register(ctx *gin.Context) {
	var u web.UserRegisterRequest
	err := ctx.BindJSON(&u)
	ok := helper.ValidationError(ctx, err)
	if ok {
		return
	}
	u.RoleId = 2
	u.VerificationTime = time.Now()
	user, err := ctrl.authService.Register(u)
	ok = helper.ValidationError(ctx, err)
	if ok {
		return
	}
	ok = helper.InternalServerError(ctx, err)
	if ok {
		return
	}
	// mainLink := helper.GetMainLink()
	// helper.SendMail(`<a href="`+mainLink+`/auth/verify_register_token/`+token+`">Click this link</a>`, "Verification Email", user.Email, user.Email, user.FirstName+" "+user.LastName)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Success",
		Errors: nil,
		Data:   user,
	}
	ctx.JSON(http.StatusCreated, webResponse)
	logger := helper.NewLog(authFile)
	logger.Infof("%d already registered", user.ID)
}

func (ctrl *authController) ForgotPasswordEmail(ctx *gin.Context) {
	var u web.UserForgotPasswordRequest
	err := ctx.BindJSON(&u)
	ok := helper.ValidationError(ctx, err)
	if ok {
		return
	}
	user := ctrl.userService.FindByEmail(u.Email)
	token, err := ctrl.jwtService.GenerateToken(strconv.FormatUint(uint64(user.ID), 10), user.Username, user.Email, uint(user.RoleId), 60*24)
	ok = helper.InternalServerError(ctx, err)
	if ok {
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
	logger := helper.NewLog(authFile)
	logger.Infof("%d already send the forgot password email", user.ID)
}

func (ctrl *authController) ForgotPassword(ctx *gin.Context) {
	var u web.UserForgotPasswordRequest
	err := ctx.BindJSON(&u)
	ok := helper.ValidationError(ctx, err)
	if ok {
		return
	}
	user := ctrl.userService.FindByUsername(u.Username)
	token, err := ctrl.jwtService.GenerateToken(strconv.FormatUint(uint64(user.ID), 10), user.Username, user.Email, uint(user.RoleId), 60*24)
	ok = helper.TokenError(ctx, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data:   token,
	}
	ctx.JSON(http.StatusOK, webResponse)
	logger := helper.NewLog(authFile)
	logger.Infof("%d already change the password", user.ID)
}

func (ctrl *authController) VerifyRegisterToken(ctx *gin.Context) {
	userToken := ctx.Param("token")
	jwtToken, err := ctrl.jwtService.ValidateToken(userToken)
	ok := helper.TokenError(ctx, err)
	if ok {
		return
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	userIdString := claims["user_id"].(string)
	userId, err := strconv.ParseUint(userIdString, 10, 64)
	helper.InternalServerError(ctx, err)
	user, err := ctrl.userService.FindById(uint(userId))
	ok = helper.NotFoundError(ctx, err)
	if ok {
		return
	}
	var userRequest web.UserRegisterVerificationTokenRequest
	userRequest.VerificationTime = time.Now()
	userRequest.ID = userId
	user, err = ctrl.authService.VerifyRegisterToken(userRequest)
	ok = helper.InternalServerError(ctx, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data:   user,
	}
	ctx.JSON(http.StatusOK, webResponse)
	logger := helper.NewLog(authFile)
	logger.Infof("%d already verify registered token", user.ID)
}

func (ctrl *authController) VerifyForgotPasswordToken(ctx *gin.Context) {
	var u web.UserNewPasswordRequest
	err := ctx.BindJSON(&u)
	ok := helper.ValidationError(ctx, err)
	if ok {
		return
	}
	userToken := ctx.Param("token")
	jwtToken, err := ctrl.jwtService.ValidateToken(userToken)
	ok = helper.TokenError(ctx, err)
	if ok {
		return
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	userIdString := claims["user_id"].(string)
	userId, err := strconv.ParseUint(userIdString, 10, 64)
	helper.InternalServerError(ctx, err)
	user, err := ctrl.userService.FindById(uint(userId))
	ok = helper.NotFoundError(ctx, err)
	if ok {
		return
	}
	u.ID = userId
	user, err = ctrl.authService.VerifyForgotPasswordToken(u)
	ok = helper.NotFoundError(ctx, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data:   user,
	}
	ctx.JSON(http.StatusOK, webResponse)
	logger := helper.NewLog(authFile)
	logger.Infof("%d already verify forgot password token", user.ID)
}
