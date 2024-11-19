package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	KafkaBrokerPORT string
	KafkaBrokerHOST string

	RedisPORT string
	RedisHOST string

	APIKey string

	AppPort string
	AppHost string
}

func LoadConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}
	
	return &Config{
		KafkaBrokerPORT: getEnv("KAFKA_BROKER_PORT", "9092"),
		KafkaBrokerHOST: getEnv("KAFKA_BROKER_HOST", "localhost"),

		RedisPORT: getEnv("REDIS_PORT", "6379"),
		RedisHOST: getEnv("REDIS_HOST", "localhost"),

		APIKey:  getEnv("API_KEY", ""),
		AppPort: getEnv("APP_PORT", "8080"),
		AppHost: getEnv("APP_HOST", "localhost"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
