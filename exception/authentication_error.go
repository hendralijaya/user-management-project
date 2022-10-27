package exception

import (
	"hendralijaya/user-management-project/helper"

	"github.com/gin-gonic/gin"
)

type AuthenticationError struct {
	context *gin.Context
	log     helper.Log
}

func NewAuthenticationError(context *gin.Context, log helper.Log) AuthenticationError {
	return AuthenticationError{
		context: context,
		log:     log,
	}
}

// this code is used to set the meta of the error
func (authError AuthenticationError) SetMeta(message error) bool {
	if(message != nil){
		authError.context.Error(message).SetMeta("UNAUTHORIZED")
		authError.context.Status(401)
		return true
	}
	return false
}

// this code is used to log the error
func (authError AuthenticationError) Logf(message error, args ...interface{}) {
	authError.log.Errorf("Authentication Error : " + message.Error(), args...)
}