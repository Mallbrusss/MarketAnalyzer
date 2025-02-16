package server

import (
	"context"
	"fmt"
	"log"
	"tinkoff-api/config"
	"tinkoff-api/internal/handlers"
	"tinkoff-api/internal/kafka"
	"tinkoff-api/internal/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	cfg         *config.Config
	e           *echo.Echo
	kafkaBroker *kafka.KafkaProducer
}

func NewServer() *Server {
	return &Server{
		cfg: config.LoadConfig(),
		e:   echo.New(),
	}
}

func (s *Server) initializeMiddleware() {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
}

func (s *Server) Shutdown(ctx context.Context) error {

	// if s.rdb != nil {
	// 	if err := s.rdb.Close(); err != nil {
	// 		log.Printf("Redis close error: %v", err)
	// 	}
	// }
	s.kafkaBroker.Close()
	return s.e.Shutdown(ctx)
}

func (s *Server) initKafka() {
	kafkaBroker := fmt.Sprintf("%s:%s", s.cfg.KafkaBrokerHOST, s.cfg.KafkaBrokerPORT)

	producer, err := kafka.NewKafkaBroker(kafkaBroker)
	if err != nil {
		log.Fatalf("Error init kafka broker: %s", err)
	}

	s.kafkaBroker = producer
}

func (s *Server) registerRoutes() {
	s.initKafka()
	// service := services.NewIngestionService(producer)

	// handlers := api.NewHandlers(service)

	// s.e.POST("/ingest-data", handlers.Ingest)

	service := services.NewTinkoffService(s.cfg, s.kafkaBroker)
	handler := handlers.NewHandler(service)

	s.e.GET("/api/v1/ti/getClosePrices", handler.GetClosePricesHandler)
	s.e.GET("/api/v1/ti/getBonds", handler.GetAllBonds)
	s.e.GET("/api/v1/ti/getCandles", handler.GetCandles)

	log.Printf("Server is running on port %s...", s.cfg.ServerPort)
}

func (s *Server) Run() error {

	// if err := s.initializeRedis(); err != nil {
	// 	s.e.Logger.Fatal("Error connecting to Redis:", err)
	// }
	// defer s.rdb.Close()

	s.initializeMiddleware()
	s.registerRoutes()

	address := fmt.Sprintf(":%s", s.cfg.ServerPort)
	return s.e.Start(address)
}
