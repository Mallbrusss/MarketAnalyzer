package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"

	"log"
)

type KafkaProducer struct {
	Producer *kafka.Producer
}

func NewKafkaBroker(broker string) (*KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  broker,
		"message.max.bytes":  5242880, // 5 МБ Fix this
		"linger.ms":          0,
		"batch.num.messages": 1,
		"group.id":           "candles",
	})
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		Producer: p,
	}, nil
}

func (kp *KafkaProducer) SendMessage(topic, key string, value []byte) error {
	log.Printf("Sending message to Kafka. Topic: %s, Key: %s, Value size: %d bytes", topic, key, len(value))
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          value,
	}

	err := kp.Producer.Produce(msg, deliveryChan)
	if err != nil {
		log.Printf("failed to send message to Kafka: %s", err)
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		log.Printf("Delivery failed: %v", m.TopicPartition.Error)
		return m.TopicPartition.Error
	}

	log.Printf("The message was successfully sent to Kafka. Topic: %s, Key: %s", topic, key)
	return nil
}

func (kp *KafkaProducer) Close() {
	kp.Producer.Close()
}
