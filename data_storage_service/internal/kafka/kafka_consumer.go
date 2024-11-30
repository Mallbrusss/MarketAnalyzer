package kafka

import (
	"data-storage/internal/models"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"log"
)

type KafkaConsumer struct {
	Consumer *kafka.Consumer

	mu           sync.Mutex
	messageParts map[string][]byte
}

func NewKafkaConsumer(broker, groupID string) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"message.max.bytes": 5242880, // 5 МБ Fix this
		"group.id":          "groupID",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		Consumer: c,
	}, nil
}

func (kc *KafkaConsumer) ListenAndProcess() {
	topics := []string{
		"candlesData",
		"bondsData",
	}
	err := kc.Consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	for {
		msg, err := kc.Consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		log.Printf("Received message from topic %s: %s", *msg.TopicPartition.Topic, msg.Value)

		switch *msg.TopicPartition.Topic {
		case "candlesData":
			if err := kc.handleCandlesData(msg); err != nil {
				log.Printf("Error processing candlesData: %v", err)
			}
		case "bondsData":
			if err := kc.handleInstruments(msg); err != nil {
				log.Printf("Error processing bondsData: %v", err)
			}
		default:
			log.Printf("No handler for topic: %s", *msg.TopicPartition.Topic)
		}
	}

}

func (kc *KafkaConsumer) handleCandlesData(msg *kafka.Message) error {
	var candles models.HistoricCandles

	err := json.Unmarshal(msg.Value, &candles)
	if err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	log.Printf("Received %d candles to process", len(candles.Candles))

	// Вставка в базу данных
	// for _, candle := range candles {
	// 	// Сохраняем каждую свечу в базу данных
	// 	err := kc.saveCandleToDB(candle)
	// 	if err != nil {
	// 		log.Printf("Failed to save candle %v: %v", candle, err)
	// 		return err
	// 	}
	// }

	return nil
}

func (kc *KafkaConsumer) handleInstruments(msg *kafka.Message) error {
	var part models.InstrumentPart

	err := json.Unmarshal(msg.Value, &part)
	if err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	log.Printf("Received part %d/%d for message ID %s", part.Part, part.Total, part.MessageID)

	kc.mu.Lock()
	defer kc.mu.Unlock()

	if _, exists := kc.messageParts[part.MessageID]; !exists {
		kc.messageParts[part.MessageID] = make([]byte, 0)
	}

	kc.messageParts[part.MessageID] = append(kc.messageParts[part.MessageID], part.Data...)

	fmt.Println(kc.messageParts)

	// Вставка в базу данных
	// for _, candle := range candles {
	// 	// Сохраняем каждую свечу в базу данных
	// 	err := kc.saveCandleToDB(candle)
	// 	if err != nil {
	// 		log.Printf("Failed to save candle %v: %v", candle, err)
	// 		return err
	// 	}
	// }

	return nil
}
