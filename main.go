package main

import (
	"hendralijaya/user-management-project/app/config"
	"hendralijaya/user-management-project/app/routes"
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.NewDB()
	mongoDB *mongo.Client = config.NewMongoDB()
)

func main() {
	defer config.CloseDB(db)
	err := godotenv.Load()
	helper.PanicIfError(err)
	router := NewRouter()
	log.Fatal(router.Run(":" + os.Getenv("GO_PORT")))
}

func NewRouter() *gin.Engine {
	err := godotenv.Load()
	helper.PanicIfError(err)
	/**
	@description Setup Database Connection
	*/

	/**
	@description Init Router
	*/
	router := gin.Default()
	/**
	@description Setup Mode Application
	*/
	if os.Getenv("GO_ENV") != "production" && os.Getenv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if os.Getenv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	/**
	@description Setup Middleware
	*/

	/**
	@description Init All Route
	*/
	routes.NewAuthenticationRoutes(db, mongoDB, router)
	routes.NewUserRoutes(db, mongoDB, router)
	routes.NewRoleRoutes(db, mongoDB, router)
	router.Use(middleware.ErrorHandler())
	router.Use(cors.Default())

	return router
}
