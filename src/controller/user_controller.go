package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nfdeveloper/crud_with_authentication/src/model/service"
)

func NewUserControllerInterface(
	serviceInteface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInteface,
	}
}

type UserControllerInterface interface {
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
