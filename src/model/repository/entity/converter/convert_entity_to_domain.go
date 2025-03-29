package converter

import (
	"github.com/nfdeveloper/crud_with_authentication/src/model"
	"github.com/nfdeveloper/crud_with_authentication/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID)
	return domain
}
