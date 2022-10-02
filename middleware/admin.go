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
	jwtService  service.JWTService
	roleService service.RoleService
}

func NewIsAdminMiddleware(jwtService service.JWTService, roleService service.RoleService) IsAdminMiddleware {
	return &isAdminMiddleware{
		jwtService:  jwtService,
		roleService: roleService,
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
		roleIdString := claims["role_id"].(string)
		roleId, err := strconv.ParseUint(roleIdString, 10, 64)
		ok = helper.InternalServerError(c, err)
		if ok {
			return
		}
		role, err := m.roleService.FindById(uint(roleId))
		ok = helper.NotFoundError(c, err)
		if ok {
			return
		}
		if role.Name != "Admin" {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: "You are not an admin",
				Data:   nil,
			}
			c.JSON(http.StatusUnauthorized, webResponse)
			c.Abort()
			return
		}
	}
}
