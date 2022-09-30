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
	TruncateTable(db, "roles")

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
	// GA PAKE / di last link
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/auth/register", requestBody)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	
}