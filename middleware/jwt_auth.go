package middleware

import (
	"hendralijaya/user-management-project/model/web"
	"hendralijaya/user-management-project/service"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthorizeJWTMiddleware interface {
	AuthorizeJWT() gin.HandlerFunc
}

type authorizeJWTMiddleware struct {
	jwtService service.JWTService
}

func NewAuthorizeJWTMiddleware(jwtService  service.JWTService) AuthorizeJWTMiddleware {
	return &authorizeJWTMiddleware{
		jwtService: jwtService,
	}
}

func (m *authorizeJWTMiddleware) AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: "Not token found",
				Data:   nil,
			}
			c.JSON(http.StatusUnauthorized, webResponse)
			c.Abort()
			return
		}
		token, err := m.jwtService.ValidateToken(authHeader)
		if err != nil {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: "Token is invalid",
				Data:   nil,
			}
			c.JSON(http.StatusUnauthorized, webResponse)
			c.Abort()
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[exp]: ", claims["name"])
			log.Println("Claim[issuer]: ", claims["issuer"])
		}
	}
}
