package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}
}

type Config struct {
	APIBaseURL string
	APIToken   string
	ServerPort string

	KafkaBrokerHOST string
	KafkaBrokerPORT string
}

func LoadConfig() *Config {

	return &Config{
		APIBaseURL: getEnv("TINKOFF_API_BASE_URL", ""),
		APIToken:   getEnv("TINKOFF_API_TOKEN", ""),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		KafkaBrokerHOST: getEnv("KAFKA_BROKER_HOST","localhost"),
		KafkaBrokerPORT: getEnv("KAFKA_BROKER_PORT","9092"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
