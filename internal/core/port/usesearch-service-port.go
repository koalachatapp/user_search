package port

import "github.com/koalachatapp/usersearch/internal/core/entity"

type UsersearchService interface {
	SearchByUUID(uuid string) error
	Save(user entity.UserEntity) error
	Delete(uuid string) error
}
