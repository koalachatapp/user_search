package main

import (
	"context"
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/koalachatapp/usersearch/cmd/kafka/handler"
	"github.com/koalachatapp/usersearch/internal/core/service"
	"github.com/koalachatapp/usersearch/internal/repository"
)

func main() {
	repo := repository.NewUsersearchRepository()
	service := service.NewUsersearchService(repo)

	log.Println("Streaming kafka")

	saramaconfig := sarama.NewConfig()
	saramaconfig.Producer.Retry.Max = 0
	saramaddr := os.Getenv("KAFKA_URL")
	if saramaddr == "" {
		saramaddr = "kafka:9092"
	}
	var cfg = sarama.NewConfig()
	c, err := sarama.NewConsumerGroup(
		[]string{saramaddr},
		"user",
		cfg,
	)
	if err != nil {
		log.Println(err)
	}
	defer c.Close()
	ctx, _ := context.WithCancel(context.Background())
	if err := c.Consume(
		ctx,
		[]string{"UsersearchTopic"},
		handler.NewKafkaHandler(service)); err != nil {
		log.Println(err)
	}

}
