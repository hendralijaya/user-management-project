package routes

import (
	"hendralijaya/user-management-project/injector"
	"hendralijaya/user-management-project/middleware"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func NewRoleRoutes(db *gorm.DB, mongoDB *mongo.Client, route *gin.Engine) {
	roleController := injector.InitRole(db)
	adminMiddleware := injector.InitAdminMiddleware(db, mongoDB)
	authorizeJWTMiddleware := injector.InitJWTMiddleware()
	// roleRoute := route.Group("/api/v1", helper.SetSession())
	roleRoute := route.Group("/api/v1/roles")
	roleRoute.Use(middleware.ErrorHandler())
	roleRoute.Use(adminMiddleware.IsAdmin())
	roleRoute.Use(authorizeJWTMiddleware.AuthorizeJWT())
	roleRoute.Use(cors.Default())
	roleRoute.GET("/", roleController.All)
	roleRoute.GET("/:id", roleController.FindById)
	roleRoute.POST("/", roleController.Insert)
	roleRoute.PUT("/:id", roleController.Update)
	roleRoute.DELETE("/:id", roleController.Delete)
}
