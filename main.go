package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nfdeveloper/crud_with_authentication/src/configuration/database/mongodb"
	"github.com/nfdeveloper/crud_with_authentication/src/configuration/logger"
	"github.com/nfdeveloper/crud_with_authentication/src/controller"
	"github.com/nfdeveloper/crud_with_authentication/src/controller/routes"
	"github.com/nfdeveloper/crud_with_authentication/src/model/service"
)

func main() {

	logger.Info("About to start user application.")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongodb.InitConnection()

	// Init Dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
