package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRoleSuccess(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)

	router := SetupRouter()
	requestBody := strings.NewReader(`{
		"name": "Test",
		"description": "desc",
	}`)
	// GA PAKE / di last link
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/roles/", requestBody)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func TestCreateRoleFailed(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)

	router := SetupRouter()
	requestBody := strings.NewReader(`{
		"name": "Test",
		"description": "desc",
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/roles/", requestBody)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}
