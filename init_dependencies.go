package main

import (
	"github.com/nfdeveloper/crud_with_authentication/src/controller"
	"github.com/nfdeveloper/crud_with_authentication/src/model/repository"
	"github.com/nfdeveloper/crud_with_authentication/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
