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

func NewUserRoutes(db *gorm.DB, route *gin.Engine) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	jwtService := service.NewJWTService()
	userController := controller.NewUserController(userService, jwtService)

	userRoute := route.Group("/api/v1/user")
	userRoute.Use(middleware.AuthorizeJWT(jwtService))
	userRoute.Use(middleware.ErrorHandler)
	userRoute.Use(cors.Default())
	userRoute.Use(middleware.IsAdmin(jwtService,userService))
	userRoute.GET("/", userController.All)
	userRoute.GET("/:id", userController.FindById)
	userRoute.POST("/", userController.Insert)
	userRoute.PUT("/:id", userController.Update)
	userRoute.DELETE("/:id", userController.Delete)
}