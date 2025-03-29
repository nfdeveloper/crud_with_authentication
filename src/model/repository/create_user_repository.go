package repository

import (
	"context"
	"os"

	"github.com/nfdeveloper/crud_with_authentication/src/configuration/logger"
	resterr "github.com/nfdeveloper/crud_with_authentication/src/configuration/rest_err"
	"github.com/nfdeveloper/crud_with_authentication/src/model"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *resterr.RestErr) {

	logger.Info("Init createUser repository")
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJsonValue()
	if err != nil {
		return nil, resterr.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, resterr.NewInternalServerError(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
