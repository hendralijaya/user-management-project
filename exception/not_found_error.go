package exception

import (
	"hendralijaya/user-management-project/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotFoundError struct {
	context *gin.Context
	log helper.Log
}

func NewNotFoundError(context *gin.Context, log helper.Log) NotFoundError {
	return NotFoundError{
		context: context,
		log: log,
	}
}

// this code is used to set the meta of the error
func (notFoundError NotFoundError) SetMeta(message error) bool {
	if (message != nil) {
		notFoundError.context.Error(message).SetMeta("NOT_FOUND_ERROR")
		notFoundError.context.Status(http.StatusNotFound)
		return true
	}
	return false
}

// this code is used to log the error
func (notFoundError NotFoundError) Logf(message error, args ...interface{}) {
	notFoundError.log.Errorf("Not Found Error : " + message.Error(), args...)
}