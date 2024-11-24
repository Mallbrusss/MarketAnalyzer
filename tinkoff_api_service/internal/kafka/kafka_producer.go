package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"

	"log"
)

type KafkaProducer struct {
	Producer *kafka.Producer
	Topic    string
}

func NewKafkaBroker(broker, topic string) (*KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		Producer: p,
		Topic:    topic,
	}, nil
}

func (kp *KafkaProducer) SendMessage(key string, value []byte) error {
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kp.Topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          value,
	}

	err := kp.Producer.Produce(msg, nil)
	if err != nil {
		log.Printf("failed to send message to Kafka: %s", err)
		return err
	}

	return nil
}

func (kp *KafkaProducer) Close() {
	kp.Producer.Close()
}
