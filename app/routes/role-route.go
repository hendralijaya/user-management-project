package routes

import (
	"hendralijaya/user-management-project/injector"
	"hendralijaya/user-management-project/middleware"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

func NewRoleRoutes(db *gorm.DB, route *gin.Engine) {
	roleController := injector.InitAdminMiddleware(db)
	// roleRoute := route.Group("/api/v1", helper.SetSession())
	roleRoute := route.Group("/api/v1/roles")
	roleRoute.Use(middleware.ErrorHandler())
	roleRoute.Use(cors.Default())
	roleRoute.GET("/", roleController.All)
	roleRoute.GET("/:id", roleController.FindById)
	roleRoute.POST("/", roleController.Insert)
	roleRoute.PUT("/:id", roleController.Update)
	roleRoute.DELETE("/:id", roleController.Delete)
}
