package repository

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/koalachatapp/usersearch/internal/core/entity"
	"github.com/koalachatapp/usersearch/internal/core/port"
)

type usersearchRepository struct {
	addr string
}

func NewUsersearchRepository() port.UsersearchRepository {
	setting, _ := sonic.Marshal(map[string]any{
		"settings": map[string]any{
			"index": map[string]any{
				"number_of_shards":   3,
				"number_of_replicas": 1,
			},
		},
	})
	usersearchrepo := usersearchRepository{
		addr: "http://localhost:9200",
	}
	// create new index if not exist
	usersearchrepo.apiCall(http.MethodPut, "/usersearchbyuuid", setting)

	return &usersearchrepo
}

func (r *usersearchRepository) apiCall(method string, path string, param []byte) map[string]any {
	req, _ := http.NewRequest(method, r.addr+path, bytes.NewBuffer(param))
	req.Header.Set("Content-Type", "application/json")
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Println(err)
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var d map[string]any
	sonic.Unmarshal(b, &d)
	return d
}

func (r *usersearchRepository) Save(user entity.UserEntity) error {
	d, _ := sonic.Marshal(user)
	resp := r.apiCall(http.MethodPost, "/usersearchbyuuid/_doc/"+user.Uuid, d)
	log.Println(resp)
	return nil
}

func (r *usersearchRepository) Delete(uuid string) error {
	resp := r.apiCall(http.MethodDelete, "/usersearchbyuuid/"+uuid, nil)
	log.Println(resp)
	return nil
}

func (r *usersearchRepository) UpdateByUUID(uuid string, user entity.UserEntity) error {
	// r.apiCall()
	return nil
}
