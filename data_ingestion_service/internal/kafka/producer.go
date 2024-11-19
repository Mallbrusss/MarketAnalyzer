package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer *kafka.Writer
}

func NewKafkaProducer(broker, topic string) *KafkaProducer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})
	return &KafkaProducer{Writer: writer}
}

func (p *KafkaProducer) SendMessage(key, value string) error{
	err := p.Writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Key: []byte(key),
			Value: []byte(value),
		},
	)
	if err != nil{
		log.Printf("Failed to write message to Kafka: %v", err)
		return err
	}
	return nil
}