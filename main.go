package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nfdeveloper/crud_with_authentication/src/configuration/database/mongodb"
	"github.com/nfdeveloper/crud_with_authentication/src/configuration/logger"
	"github.com/nfdeveloper/crud_with_authentication/src/controller/routes"
)

func main() {

	logger.Info("About to start user application.")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database, err := mongodb.NewMongoDBConnection(context.Background())

	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
