package producer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer interface {
	Produce(topic string, message string) error
}

type producerImpl struct {
	p *kafka.Producer
}

func NewProducer(producer *kafka.Producer) Producer {
	return &producerImpl{p: producer}
}

func (producer *producerImpl) Produce(topic string, message string) error {
	deliveryChan := make(chan kafka.Event)
	err := producer.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)

	if err != nil {
		fmt.Printf("Produce failed: %v\n", err)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
	return m.TopicPartition.Error
}
