package service

import (
	"github.com/koalachatapp/usersearch/internal/core/entity"
	"github.com/koalachatapp/usersearch/internal/core/port"
)

type usersearchService struct {
	repo port.UsersearchRepository
}

func NewUsersearchService(repo port.UsersearchRepository) port.UsersearchService {

	return &usersearchService{
		repo: repo,
	}
}

func (s *usersearchService) SearchByUUID(uuid string) error {
	return nil
}

func (s *usersearchService) Save(user entity.UserEntity) error {
	return s.repo.Save(user)
}

func (s *usersearchService) Delete(uuid string) error {
	return s.repo.Delete(uuid)
}
