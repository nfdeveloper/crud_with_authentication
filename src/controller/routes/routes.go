package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nfdeveloper/crud_with_authentication/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

}
