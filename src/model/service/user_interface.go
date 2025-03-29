package service

import (
	resterr "github.com/nfdeveloper/crud_with_authentication/src/configuration/rest_err"
	"github.com/nfdeveloper/crud_with_authentication/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *resterr.RestErr
	UpdateUser(string, model.UserDomainInterface) *resterr.RestErr
	FindUser(string) (*model.UserDomainInterface, *resterr.RestErr)
	DeleteUser(string) *resterr.RestErr
}
