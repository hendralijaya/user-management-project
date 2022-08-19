package controller

import (
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/model/domain"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/service"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Logout(ctx *gin.Context)
	ForgotPassword(ctx *gin.Context)
	VerifyRegisterToken(ctx *gin.Context)
	VerifyForgotPasswordToken(ctx *gin.Context)
}

type authController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewAuthController(userService service.UserService, jwtService service.JWTService) AuthController {
	return &authController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var u web.UserLoginRequest
	err := ctx.BindJSON(&u)
	ok := helper.ValidationError(ctx, err)
	if ok {
		return
	}
	user, err := c.userService.VerifyCredential(u)
	ok = helper.AuthenticationError(ctx, err)
	if ok {
		return
	}
	if v, ok := user.(domain.User); ok {
		generateToken, err := c.jwtService.GenerateToken(strconv.FormatUint(v.Id, 10), v.Name)
		ok = helper.InternalServerError(ctx, err)
		if ok {
			return
		}
		if v.VerificationTime.IsZero() {
			webResponse := web.WebResponse{
				Code:   http.StatusOK,
				Status: "Success",
				Errors: "Please verify your account",
				Data:   nil,
			}
			ctx.JSON(http.StatusOK, webResponse)
		}
		v.Token = generateToken
		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: "Success",
			Errors: nil,
			Data:   v,
		}
		ctx.JSON(http.StatusOK, webResponse)
		return
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var u web.UserRegisterRequest
	err := ctx.BindJSON(&u)
	ok := helper.ValidationError(ctx, err)
	if ok {
		return
	}
	u.Role_id = 2
	user, err := c.userService.Create(u)
	ok = helper.ValidationError(ctx, err)
	if ok {
		return
	}
	userIdString := strconv.FormatUint(user.Id, 10)
	token, err := service.JWTService.GenerateToken(c.jwtService, userIdString, user.Name)
	ok = helper.InternalServerError(ctx, err)
	if ok {
		return
	}
	mainLink := helper.GetMainLink()
	helper.SendMail(`<a href="`+mainLink+`/verify_register_token/`+token+`">Click this link</a>`, "Verification Email", user.Email, user.Email, user.Name)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Success",
		Errors: nil,
		Data:   user,
	}
	ctx.JSON(http.StatusCreated, webResponse)
}

func (c *authController) Logout(ctx *gin.Context) {
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *authController) ForgotPassword(ctx *gin.Context) {
	var u web.UserForgotPasswordRequest
	err := ctx.BindJSON(&u)
	ok := helper.ValidationError(ctx, err)
	if ok {
		return
	}
	user, err := c.userService.FindByEmail(u.Email)
	ok = helper.NotFoundError(ctx, err)
	if ok {
		return
	}
	token, err := c.jwtService.GenerateToken(strconv.FormatUint(user.Id, 10), user.Name)
	ok = helper.InternalServerError(ctx, err)
	if ok {
		return
	}
	mainLink := helper.GetMainLink()
	helper.SendMail(`<a href="`+mainLink+`/verify_forgot_password_token/`+token+`">Click this link</a>`, "Forgot Password Email", user.Email, user.Email, user.Name)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data:   user,
	}
	ctx.JSON(http.StatusOK, webResponse)

}

func (c *authController) VerifyRegisterToken(ctx *gin.Context) {
	userToken := ctx.Param("token")
	jwtToken, err := c.jwtService.ValidateToken(userToken)
	helper.TokenError(ctx, err)
	claims := jwtToken.Claims.(jwt.MapClaims)
	userIdString := claims["user_id"]
	userId, err := strconv.ParseUint(userIdString.(string), 10, 64)
	helper.InternalServerError(ctx, err)
	user, err := c.userService.FindById(userId)
	ok := helper.NotFoundError(ctx, err)
	if ok {
		return
	}
	user.VerificationTime = time.Now()
	user, err = c.userService.Update(user)
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
}

func (c *authController) VerifyForgotPasswordToken(ctx *gin.Context) {
	var u web.UserNewPasswordRequest
	err := ctx.BindJSON(&u)
	ok := helper.ValidationError(ctx, err)
	if ok {
		return
	}
	userToken := ctx.Param("token")
	jwtToken, err := c.jwtService.ValidateToken(userToken)
	helper.TokenError(ctx, err)
	claims := jwtToken.Claims.(jwt.MapClaims)
	userIdString := claims["user_id"]
	userId, err := strconv.ParseUint(userIdString.(string), 10, 64)
	helper.InternalServerError(ctx, err)
	user, err := c.userService.FindById(userId)
	ok = helper.NotFoundError(ctx, err)
	if ok {
		return
	}
	user.Password = u.Password
	user, err = c.userService.Update(user)
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
}
