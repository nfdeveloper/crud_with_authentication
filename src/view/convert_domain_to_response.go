package view

import (
	"github.com/nfdeveloper/crud_with_authentication/src/controller/model/response"
	"github.com/nfdeveloper/crud_with_authentication/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
