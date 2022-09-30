package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUserSuccess(t *testing.T) {
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

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/auth/register", requestBody)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusCreated, response.StatusCode)

}

func TestRegisterUserFailed(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)

	router := SetupRouter()
	requestBody := strings.NewReader(`{
		"username": "Testimonial",
 		"last_name": "Last name",
 		"nik": "123456123",
 		"address": "blablabla",
		"phone_number": "081234567891",
 		"gender": "Male",
 		"email": "coba@gmail.com",
 		"password": "12345678",
 		"created_by": "Testimonial"
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/auth/register", requestBody)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

}

func TestLoginUserSuccess(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)
	authHeader := LoginGet()

	router := SetupRouter()
	requestBody := strings.NewReader(`{
		"username": "Testimonial",
		"email": "coba@gmail.com",
		"password": "12345678"
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/auth/login", requestBody)
	request.Header.Add("Authorization", authHeader)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestLoginUserFailedUnauthorized(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)

	router := SetupRouter()
	requestBody := strings.NewReader(`{
		"username": "Testimonial",
		"email": "coba@gmail.com",
		"password": "12345678"
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/auth/login", requestBody)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestForgotPasswordSuccess(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)

	router := SetupRouter()
	requestBody := strings.NewReader(`{
		"email": "coba@gmail.com",
		"repeat_password": "11111111",
		"password": "11111111"
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/auth/forgot_password", requestBody)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestForgotPasswordFailed(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)

	router := SetupRouter()
	requestBody := strings.NewReader(`{
		"email": "coba@gmail.com",
		"repeat_password": "11111111",
		"password": "11111111"
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/auth/forgot_password", requestBody)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
