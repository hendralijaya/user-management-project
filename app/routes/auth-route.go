package routes

import (
	"hendralijaya/user-management-project/controller"
	"hendralijaya/user-management-project/middleware"
	"hendralijaya/user-management-project/repository"
	"hendralijaya/user-management-project/service"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

func NewAuthenticationRoutes(db *gorm.DB, route *gin.Engine) {
	userRepository := repository.NewUserRepository(db)
	authService := service.NewUserService(userRepository)
	jwtService := service.NewJWTService()
	authController := controller.NewAuthController(authService, jwtService)

	// authRoute := route.Group("/api/v1", helper.SetSession())
	authRoute := route.Group("/api/v1")
	authRoute.Use(middleware.ErrorHandler)
	authRoute.Use(cors.Default())
	authRoute.POST("/login/", authController.Login)
	authRoute.POST("/register/", authController.Register)
	authRoute.POST("/logout/", authController.Logout)
	authRoute.POST("/forgot_password/", authController.ForgotPassword)
	authRoute.POST("/verify_register_token/:token", authController.VerifyRegisterToken)
}
