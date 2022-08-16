package middleware

import (
	"hendralijaya/austin-hendra-restapi/model/web"
	"hendralijaya/austin-hendra-restapi/service"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			webResponse := web.WebResponse{
				Code:  http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: "Not token found",
				Data: nil,
			}
			c.JSON(http.StatusUnauthorized, webResponse)
			return
		}
		token, _ := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[exp]: ", claims["username"])
			log.Println("Claim[issuer]: ", claims["issuer"])
		}else {
			webResponse := web.WebResponse{
				Code:  http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: "Token is invalid",
				Data: nil,
			}
			c.JSON(http.StatusUnauthorized, webResponse)
		}
	}
}