package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	resterr "github.com/nfdeveloper/crud_with_authentication/src/configuration/rest_err"
	"github.com/nfdeveloper/crud_with_authentication/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := resterr.NewBadRequestError(
			fmt.Sprintf("There are some incorret fields, erro=%s ", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
