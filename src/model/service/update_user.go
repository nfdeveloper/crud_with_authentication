package service

import (
	resterr "github.com/nfdeveloper/crud_with_authentication/src/configuration/rest_err"
	"github.com/nfdeveloper/crud_with_authentication/src/model"
)

func (ud *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *resterr.RestErr {
	return nil
}
