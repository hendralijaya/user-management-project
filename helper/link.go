package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func GetMainLink() string {
	err := godotenv.Load()
	PanicIfError(err)

	return "http://localhost:" + os.Getenv("PORT") + "/api/v1"
}