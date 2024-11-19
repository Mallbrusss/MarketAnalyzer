package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	KafkaBrokerPORT  string
	KafkaBrokerHOST string

	RedisPORT  string
	RedisHOST string

	APIKey string

	AppPort string
}

func LoadConfig() *Config {
	godotenv.Load()
	return &Config{
		KafkaBrokerPORT: getEnv("KAFKA_BROKER_PORT","9092"),
		KafkaBrokerHOST: getEnv("KAFKA_BROKER_HOST","localhost"),

		RedisPORT: getEnv("REDIS_PORT","6379"),
		RedisHOST: getEnv("REDIS_HOST","localhost"),

		APIKey: getEnv("API_KEY",""),
		AppPort: getEnv("APP_PORT","8080"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
