package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"golang-echo-kafka-example/model"
)

type Consumer interface {
	Consume(topics []string)
}

type consumerImpl struct {
	c *kafka.Consumer
}

func NewConsumer(consumer *kafka.Consumer) Consumer {
	return &consumerImpl{c: consumer}
}

func (consumer *consumerImpl) Consume(topics []string) {
	consumer.c.SubscribeTopics(topics, nil)
	for {
		msg, err := consumer.c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			var user model.User
			err = json.Unmarshal(msg.Value, &user)
			if err != nil {
				fmt.Printf("Json processing error: %s\n", err.Error())
			}

			fmt.Printf("received message from event : %s\n", user)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
