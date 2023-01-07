package port

import "github.com/koalachatapp/usersearch/internal/core/entity"

type UsersearchRepository interface {
	Save(user entity.UserEntity) error
	Delete(user string) error
}
