package middleware

import (
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
	jwtService service.JWTService
	userService service.UserService
}

func NewIsAdminMiddleware(jwtService service.JWTService, userService service.UserService) IsAdminMiddleware {
	return &isAdminMiddleware{
		jwtService: jwtService,
		userService: userService,
	}
}

func (m *isAdminMiddleware) IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		jwtToken, err := m.jwtService.ValidateToken(token)
		ok := helper.TokenError(c, err)
		if ok {
			return
		}
		claims := jwtToken.Claims.(jwt.MapClaims)
		userIdString := claims["user_id"].(string)
		userId, err := strconv.ParseUint(userIdString, 10, 64)
		ok = helper.InternalServerError(c, err)
		if ok {
			return
		}
		user, err := m.userService.FindById(userId)
		ok = helper.NotFoundError(c, err)
		if ok {
			return
		}
		if(user.Role_id != 1) {
			webResponse := web.WebResponse{
				Code:  http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: "You are not an admin",
				Data: nil,
			}
			c.JSON(http.StatusUnauthorized, webResponse)
			c.Abort()
			return
		}
		c.Next()
	}
}