package converter

import (
	"github.com/nfdeveloper/crud_with_authentication/src/model"
	"github.com/nfdeveloper/crud_with_authentication/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
