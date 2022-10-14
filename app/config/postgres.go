package config

import (
	"hendralijaya/user-management-project/helper"
	"hendralijaya/user-management-project/model/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	err := godotenv.Load()
	helper.PanicIfError(err)

	// dbUser := os.Getenv("POSTGRES_USER")
	// dbPass := os.Getenv("POSTGRES_PASSWORD")
	// dbHost := os.Getenv("POSTGRES_HOST")
	// dbPort := os.Getenv("POSTGRES_PORT")
	// dbName := os.Getenv("POSTGRES_NAME")
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	dbURL := "postgres://root:root@172.31.1.3:5432/user-management?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	helper.PanicIfError(err)
	db.AutoMigrate(&domain.User{}, &domain.Role{})
	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	helper.PanicIfError(err)
	dbSQL.Close()
}
