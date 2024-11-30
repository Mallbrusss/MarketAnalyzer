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

	PostgresPort     string
	PostgresHost     string
	PostgresPassword string
	PostgresUser     string
	PostgresDatabase string

	KafkaBrokerHOST string
	KafkaBrokerPORT string
}

func LoadConfig() *Config {

	return &Config{
		APIBaseURL:      getEnv("TINKOFF_API_BASE_URL", ""),
		APIToken:        getEnv("TINKOFF_API_TOKEN", ""),
		ServerPort:      getEnv("SERVER_PORT", "8081"),
		KafkaBrokerHOST: getEnv("KAFKA_BROKER_HOST", "localhost"),
		KafkaBrokerPORT: getEnv("KAFKA_BROKER_PORT", "9092"),

		PostgresPort:     getEnv("POSTGRES_PORT", "8080"),
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "testpassword"),
		PostgresUser:     getEnv("POSTGRES_USER", "timescale"),
		PostgresDatabase: getEnv("POSTGRES_DATABASE", "timescale"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
