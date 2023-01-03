package handler

import (
	"log"

	"github.com/Shopify/sarama"
	"github.com/bytedance/sonic"
	"github.com/koalachatapp/usersearch/internal/core/entity"
)

type KafkaHandler struct{}

func NewKafkaHandler() *KafkaHandler {
	return &KafkaHandler{}
}

func (KafkaHandler) Cleanup(g sarama.ConsumerGroupSession) error {
	log.Println("Ended connection")
	return nil
}

func (KafkaHandler) Setup(g sarama.ConsumerGroupSession) error {
	log.Println("Prepare connection")
	return nil
}

func (KafkaHandler) ConsumeClaim(g sarama.ConsumerGroupSession, c sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg := <-c.Messages():
			var data entity.UserEventEntity
			_ = sonic.Unmarshal(msg.Value, &data)
			switch data.Method {
			case "register":
				log.Println(data.Data)
			}
			g.MarkMessage(msg, "Done")
		case <-g.Context().Done():
			return nil
		}
	}
}
