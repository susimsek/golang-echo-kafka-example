package config

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

var (
	UserNotificationTopic = "user_notification_event"
	UserConsumerGroup     = "userGroup"
)

func InitKafkaProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": BootstrapServers})
	if err != nil {
		log.Fatal(err)
	}

	return p
}

func InitKafkaConsumer(groupId string) *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": BootstrapServers,
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatal(err)
	}

	return c
}
