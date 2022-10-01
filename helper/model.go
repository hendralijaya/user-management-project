package helper

import (
	"hendralijaya/user-management-project/model/web"

	"github.com/mashingan/smapping"
)

func ConvertToUserUpdateRequest(i interface{}) (web.UserUpdateRequest, error) {
	var userRequest web.UserUpdateRequest
	err := smapping.FillStruct(userRequest, smapping.MapFields(i))
	return userRequest, err
}