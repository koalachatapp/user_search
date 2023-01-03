package port

import "github.com/koalachatapp/usersearch/internal/core/entity"

type UsersearchService interface {
	Search(user entity.UserEntity) error
	Save(user entity.UserEntity) error
}
