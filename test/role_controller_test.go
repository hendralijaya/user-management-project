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
	TruncateTable(db, "users")

	router := SetupRouter()
	requestBody := strings.NewReader(`{
		
	}`)
	// GA PAKE / di last link
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/auth/register", requestBody)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusCreated, response.StatusCode)
	
}