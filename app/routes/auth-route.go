package routes

import (
	"hendralijaya/user-management-project/injector"
	"hendralijaya/user-management-project/middleware"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

func NewAuthenticationRoutes(db *gorm.DB, route *gin.Engine) {
	authController := injector.InitAuth(db)
	// authRoute := route.Group("/api/v1", helper.SetSession())
	authRoute := route.Group("/api/v1/auth")
	authRoute.Use(middleware.ErrorHandler())
	authRoute.Use(cors.Default())
	authRoute.POST("/login", authController.Login)
	authRoute.POST("/register", authController.Register)
	authRoute.POST("/forgot-password", authController.ForgotPassword)
	authRoute.POST("/forgot_password_email", authController.ForgotPasswordEmail)
	authRoute.POST("/verify_register_token/:token", authController.VerifyRegisterToken)
	authRoute.POST("/verify_forgot_password_token/:token", authController.VerifyForgotPasswordToken)
}
