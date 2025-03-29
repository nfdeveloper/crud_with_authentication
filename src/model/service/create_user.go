package service

import (
	"fmt"

	"github.com/nfdeveloper/crud_with_authentication/src/configuration/logger"
	resterr "github.com/nfdeveloper/crud_with_authentication/src/configuration/rest_err"
	"github.com/nfdeveloper/crud_with_authentication/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *resterr.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(ud)

	return nil
}
