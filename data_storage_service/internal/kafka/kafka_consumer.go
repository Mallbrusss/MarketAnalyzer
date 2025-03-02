package kafka

import (
	"bytes"
	"data-storage/internal/models"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"log"
)

type InstrumentRepositoryInterface interface {
	CreateInstruments(instruments []models.PlacementPrice) error
	CreateCandles(instruments []models.HistoricCandle) error
}

type KafkaConsumer struct {
	Consumer                                  *kafka.Consumer
	mu                                        sync.Mutex
	recieveInstrumentPart, recieveCandlesPart []models.InstrumentPart
	ir                                        InstrumentRepositoryInterface
}

// type MessageAssembly struct {
// 	TotalParts int
// 	Received   int
// 	Data       []byte
// }

func NewKafkaConsumer(broker, groupID string, ir InstrumentRepositoryInterface) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"message.max.bytes": 5242880, // 5 МБ Fix this
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		Consumer:              c,
		ir:                    ir,
		recieveInstrumentPart: make([]models.InstrumentPart, 0),
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

		// log.Printf("Received message from topic %s: %s", *msg.TopicPartition.Topic, msg.Value)

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
	var part models.InstrumentPart

	err := json.Unmarshal(msg.Value, &part)
	if err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	log.Printf("Received part %d/%d for message ID %s", part.Part, part.Total, part.MessageID)

	kc.mu.Lock()
	defer kc.mu.Unlock()

	kc.recieveCandlesPart = append(kc.recieveCandlesPart, part)
	if len(kc.recieveCandlesPart) == 0 {
		log.Println("Recieve message is nil")
	}

	if kc.isMessageComplete(kc.recieveCandlesPart, part.Total) {
		data, err := kc.buildAllData(kc.recieveCandlesPart)
		if err != nil {
			log.Printf("failed to build complete message: %s", err)
			return err
		}

		if err := kc.processCompleteMessageCandles(data); err != nil {
			log.Printf("failed to process complete message: %s", err)
			return err
		}
	}

	return nil
}

func (kc *KafkaConsumer) processCompleteMessageCandles(msg []byte) error {
	log.Println("Processing complete message Candles")

	var candles models.HistoricCandles
	if err := json.Unmarshal(msg, &candles); err != nil {
		log.Printf("failed to unmarshal complete message: %s", err)
		return err
	}

	if err := kc.ir.CreateCandles(candles.Candles); err != nil {
		log.Printf("failed to save message to database: %s", err)
		return err
	}

	log.Println("Message successfully processed and saved")
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

	kc.recieveInstrumentPart = append(kc.recieveInstrumentPart, part)

	if kc.isMessageComplete(kc.recieveInstrumentPart, part.Total) {
		data, err := kc.buildAllData(kc.recieveInstrumentPart)
		if err != nil {
			log.Printf("failed to build complete message: %s", err)
			return err
		}

		if err := kc.processCompleteMessageInstruments(data); err != nil {
			log.Printf("failed to process complete message: %s", err)
			return err
		}
	}

	return nil
}

func (kc *KafkaConsumer) buildAllData(ip []models.InstrumentPart) ([]byte, error) {
	var allData bytes.Buffer

	sort.Slice(ip, func(i, j int) bool {
		return ip[i].Part < ip[j].Part
	})

	for _, part := range ip {
		_, err := allData.Write(part.Data)
		if err != nil {
			log.Println("Error build all data")
			return nil, err
		}
	}

	return allData.Bytes(), nil
}

func (kc *KafkaConsumer) processCompleteMessageInstruments(msg []byte) error {
	log.Println("Processing complete message Instruments")

	var instruments models.Instruments
	if err := json.Unmarshal(msg, &instruments); err != nil {
		return fmt.Errorf("failed to unmarshal complete message: %w", err)
	}

	if err := kc.ir.CreateInstruments(instruments.Instruments); err != nil {
		return fmt.Errorf("failed to save message to database: %w", err)
	}

	log.Println("Message successfully processed and saved")
	return nil
}

func (kc *KafkaConsumer) processCompleteMessageBonds(msg []byte) error {
	log.Println("Processing complete message Instruments")

	var candles []models.HistoricCandle
	if err := json.Unmarshal(msg, &candles); err != nil {
		return fmt.Errorf("failed to unmarshal complete message: %w", err)
	}

	if err := kc.ir.CreateCandles(candles); err != nil {
		return fmt.Errorf("failed to save message to database: %w", err)
	}

	log.Println("Message successfully processed and saved")
	return nil
}

func (kc *KafkaConsumer) isMessageComplete(ip []models.InstrumentPart, totalParts int) bool {
	if len(ip) != totalParts {
		fmt.Println("false")
		return false
	}

	fmt.Println("true")
	return true
}
