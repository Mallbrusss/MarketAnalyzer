package server

import (
	"context"
	"data-storage/config"
	"data-storage/internal/kafka"
	"data-storage/internal/storage/timescale"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// "github.com/redis/go-redis"
	"gorm.io/gorm"
)

type Server struct {
	cfg           *config.Config
	e             *echo.Echo
	db            *gorm.DB
	kafkaConsumer *kafka.KafkaConsumer
	// rdb       *redis.Client
}

func NewServer() *Server {
	return &Server{
		cfg: config.LoadConfig(),
		e:   echo.New(),
	}
}

func (s *Server) initializeDatabase() error {
	db, err := timescale.InitDB(s.cfg.PostgresHost, s.cfg.PostgresUser,
		s.cfg.PostgresPassword, s.cfg.PostgresDatabase, s.cfg.PostgresPort)
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Server) initializeRedis() error {
	// rdb, err := inRdb.InitRedisCl()
	// if err != nil {
	// 	return err
	// }
	// s.rdb = rdb
	return nil
}

func (s *Server) initializeMiddleware() {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
}

func (s *Server) initKafka() {
	kafkaBroker := fmt.Sprintf("%s:%s", s.cfg.KafkaBrokerHOST, s.cfg.KafkaBrokerPORT)
	groupID := "dataProcessorGroup"

	consumer, err := kafka.NewKafkaConsumer(kafkaBroker, groupID)
	if err != nil {
		log.Fatalf("Error init kafka broker: %s", err)
	}

	s.kafkaConsumer = consumer
}

func (s *Server) startKafka() {
	go s.kafkaConsumer.ListenAndProcess()
}

func (s *Server) registerRoutes() {
	
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}

func (s *Server) Run() error {
	s.initializeDatabase()

	s.initKafka()
	s.startKafka()

	s.initializeMiddleware()
	s.registerRoutes()

	address := fmt.Sprintf(":%s", s.cfg.ServerPort)
	return s.e.Start(address)
}
