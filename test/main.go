package test

import (
	"fmt"
	"hendralijaya/user-management-project/app/routes"
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/middleware"
	"hendralijaya/user-management-project/model/domain"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = setupTestDB()
)

func setupTestDB() *gorm.DB {
	err := godotenv.Load()
	helper.PanicIfError(err)

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	db.AutoMigrate(&domain.User{}, &domain.Role{})
	return db
}

func CloseTestDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	helper.PanicIfError(err)
	dbSQL.Close()
}

func TruncateTable(db *gorm.DB, tableName string) {
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", tableName))
}

func SetupRouter() *gin.Engine {
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
	routes.NewAuthenticationRoutes(db, router)
	router.Use(middleware.ErrorHandler())
	router.Use(cors.Default())

	return router
}

func LoginGet() string{
	db := setupTestDB()
	defer CloseTestDB(db)
	TruncateTable(db, "users")

	router := SetupRouter()
	requestBody := strings.NewReader(`{
		"username": "Testimonial",
 		"first_name": "First name",
 		"last_name": "Last name",
 		"nik": "123456123",
 		"address": "blablabla",
		"phone_number": "081234567891",
 		"gender": "Male",
 		"email": "coba@gmail.com",
 		"password": "12345678",
 		"created_by": "Testimonial"
	}`)
	// GA PAKE / di last link
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/auth/register", requestBody)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	return response.Header.Get("Authorization")
}