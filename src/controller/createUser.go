package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nfdeveloper/crud_with_authentication/src/configuration/logger"
	"github.com/nfdeveloper/crud_with_authentication/src/configuration/validation"
	"github.com/nfdeveloper/crud_with_authentication/src/controller/model/request"
	"github.com/nfdeveloper/crud_with_authentication/src/model"
	"github.com/nfdeveloper/crud_with_authentication/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domain, err := uc.service.CreateUser(domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domain,
	))
}
