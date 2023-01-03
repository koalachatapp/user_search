package repository

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/koalachatapp/usersearch/internal/core/entity"
	"github.com/koalachatapp/usersearch/internal/core/port"
)

type usersearchRepository struct {
	client *elasticsearch.Client
}

func NewUsersearchRepository() port.UsersearchRepository {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Println(err)
	}
	return &usersearchRepository{
		client: client,
	}
}

func (r *usersearchRepository) Save(user entity.UserEntity) error {
	return nil
}
