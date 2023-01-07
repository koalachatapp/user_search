package handler

import (
	"log"

	"github.com/Shopify/sarama"
	"github.com/bytedance/sonic"
	"github.com/koalachatapp/usersearch/internal/core/entity"
	"github.com/koalachatapp/usersearch/internal/core/port"
)

type KafkaHandler struct {
	service port.UsersearchService
}

func NewKafkaHandler(service port.UsersearchService) *KafkaHandler {
	return &KafkaHandler{
		service: service,
	}
}

func (k *KafkaHandler) Cleanup(g sarama.ConsumerGroupSession) error {
	log.Println("Ended connection")
	return nil
}

func (k *KafkaHandler) Setup(g sarama.ConsumerGroupSession) error {
	log.Println("Prepare connection")
	return nil
}

func (k *KafkaHandler) ConsumeClaim(g sarama.ConsumerGroupSession, c sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg := <-c.Messages():
			var data entity.UserEventEntity
			_ = sonic.Unmarshal(msg.Value, &data)
			switch data.Method {
			case "register":
				k.Register(data.Data)
			case "update":
				k.Update(data.Data)
			case "delete":
				k.Delete(data.Data)
			}
			g.MarkMessage(msg, "Done")
		case <-g.Context().Done():
			return nil
		}
	}
}

func (k *KafkaHandler) Register(user entity.UserEntity) {
	k.service.Save(user)
}

func (k *KafkaHandler) Update(user entity.UserEntity) {

}

func (k *KafkaHandler) Delete(user entity.UserEntity) {
	k.service.Delete(user.Uuid)
}
