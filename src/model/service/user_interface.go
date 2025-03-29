package service

import (
	resterr "github.com/nfdeveloper/crud_with_authentication/src/configuration/rest_err"
	"github.com/nfdeveloper/crud_with_authentication/src/model"
	"github.com/nfdeveloper/crud_with_authentication/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository: userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *resterr.RestErr)
	UpdateUser(string, model.UserDomainInterface) *resterr.RestErr
	FindUser(string) (*model.UserDomainInterface, *resterr.RestErr)
	DeleteUser(string) *resterr.RestErr
}
