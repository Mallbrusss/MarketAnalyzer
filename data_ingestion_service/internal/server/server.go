package server

import (
	"context"
	"data_ingestion/config"
	"data_ingestion/internal/api"
	"data_ingestion/internal/kafka"
	"data_ingestion/internal/services"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/redis/go-redis"
)

type Server struct {
	cfg *config.Config
	e   *echo.Echo
	// rdb       *redis.Client
}

func NewServer() *Server {
	return &Server{
		cfg: config.LoadConfig(),
		e:   echo.New(),
	}
}

// func (s *Server) initializeRedis() error {
// 	rdb, err := inRdb.InitRedisCl()
// 	if err != nil {
// 		return err
// 	}
// 	s.rdb = rdb
// 	return nil
// }

func (s *Server) initializeMiddleware() {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
}

func (s *Server) registerRoutes() {
	kafkaBroker := fmt.Sprintf("%s:%s", s.cfg.KafkaBrokerHOST, s.cfg.KafkaBrokerPORT)

	producer := kafka.NewKafkaProducer(kafkaBroker, "data-ingestion")

	service := services.NewIngestionService(producer)

	handlers := api.NewHandlers(service)

	s.e.POST("/ingest-data", handlers.Ingest)
}

func (s *Server) Shutdown(ctx context.Context) error {

	// if s.rdb != nil {
	// 	if err := s.rdb.Close(); err != nil {
	// 		log.Printf("Redis close error: %v", err)
	// 	}
	// }
	return s.e.Shutdown(ctx)
}

func (s *Server) Run() error {

	// if err := s.initializeRedis(); err != nil {
	// 	s.e.Logger.Fatal("Error connecting to Redis:", err)
	// }
	// defer s.rdb.Close()

	s.initializeMiddleware()
	s.registerRoutes()

	address := fmt.Sprintf("%s:%s", s.cfg.AppHost, s.cfg.AppPort)
	return s.e.Start(address)
}
