package routes

import (
	"hendralijaya/user-management-project/injector"
	"hendralijaya/user-management-project/middleware"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

func NewUserRoutes(db *gorm.DB, route *gin.Engine) {
	userController := injector.InitUser(db)
	authorizeJWTMiddleware := injector.InitJWTMiddleware()
	adminMiddleware := injector.InitAdminMiddleware(db)
	userRoute := route.Group("/api/v1/user")
	userRoute.Use(authorizeJWTMiddleware.AuthorizeJWT())
	userRoute.Use(middleware.ErrorHandler())
	userRoute.Use(cors.Default())
	userRoute.Use(adminMiddleware.IsAdmin())
	userRoute.GET("/", userController.All)
	userRoute.GET("/:id", userController.FindById)
	userRoute.POST("/", userController.Insert)
	userRoute.PUT("/:id", userController.Update)
	userRoute.DELETE("/:id", userController.Delete)
}