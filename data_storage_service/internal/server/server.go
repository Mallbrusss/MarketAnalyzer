package server

import (
	"context"
	"data-storage/config"
	"data-storage/internal/kafka"
	"data-storage/internal/repository"
	"data-storage/internal/storage/timescale"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		log.Fatal(err)
		return err
	}

	s.db = db
	return nil
}

func (s *Server) initializeRepository(db *gorm.DB) *repository.InstrumentRepository {
	return repository.NewInstrumentRepository(db)
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

func (s *Server) initializeKafka(repository *repository.InstrumentRepository) {
	kafkaBroker := fmt.Sprintf("%s:%s", s.cfg.KafkaBrokerHOST, s.cfg.KafkaBrokerPORT)
	groupID := "dataProcessorGroup"

	consumer, err := kafka.NewKafkaConsumer(kafkaBroker, groupID, repository)
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

func (s *Server) downloadInstruments() error {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/getBonds", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("non-200 status code received")
	}
	return nil
}

func (s *Server) Run() error {
	log.Println("Init database")
	err := s.initializeDatabase()
	if err != nil {
		log.Println("Database initialization error")
	}
	log.Println("Database init success")

	repository := s.initializeRepository(s.db)

	s.initializeKafka(repository)
	s.startKafka()

	s.initializeMiddleware()
	s.registerRoutes()
	err = s.downloadInstruments()
	if err != nil {
		log.Printf("can`t download instruments: %s", err)
	}

	address := fmt.Sprintf(":%s", s.cfg.ServerPort)
	return s.e.Start(address)
}
