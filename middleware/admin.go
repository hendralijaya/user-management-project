package middleware

import (
	"hendralijaya/user-management-project/exception"
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type IsAdminMiddleware interface {
	IsAdmin() gin.HandlerFunc
}

type isAdminMiddleware struct {
	jwtService  service.JWTService
	userService service.UserService
	logger 	helper.Log
}

func NewIsAdminMiddleware(jwtService service.JWTService, userService service.UserService, logger helper.Log) IsAdminMiddleware {
	return &isAdminMiddleware{
		jwtService:  jwtService,
		userService: userService,
		logger: logger,
	}
}

func (middleware *isAdminMiddleware) IsAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenError := exception.NewTokenError(context, middleware.logger)
		internalServerError := exception.NewInternalServerError(context, middleware.logger)
		notFoundError := exception.NewNotFoundError(context, middleware.logger)
		token := context.GetHeader("Authorization")
		jwtToken, err := middleware.jwtService.ValidateToken(token)
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
		user, err := middleware.userService.FindById(uint(userId))
		ok = notFoundError.SetMeta(err)
		if ok {
			notFoundError.Logf(err)
			return
		}
		if user.RoleId != 1 {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: "You are not an admin",
				Data:   nil,
			}
			context.JSON(http.StatusUnauthorized, webResponse)
			context.Abort()
			return
		}
	}
}
