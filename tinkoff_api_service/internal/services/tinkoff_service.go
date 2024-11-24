package services

import (
	"encoding/json"
	"fmt"
	"time"
	"tinkoff-api/config"
	"tinkoff-api/internal/kafka"
	"tinkoff-api/internal/models"
	"tinkoff-api/pkg/httpclient"
)

type TinkoffService struct {
	Client        *httpclient.HTTPClient
	Config        *config.Config
	KafkaProducer *kafka.KafkaProducer
}

func NewTinkoffService(cfg *config.Config, kafkaProducer *kafka.KafkaProducer) *TinkoffService {
	return &TinkoffService{
		Client:        httpclient.NewHTTPClient(),
		Config:        cfg,
		KafkaProducer: kafkaProducer,
	}
}

func (s *TinkoffService) GetClosePrices(instruments []string) ([]models.ClosePrice, error) {
	var InstrumentRequests []models.InstrumentRequest
	for _, instrument := range instruments {
		InstrumentRequests = append(InstrumentRequests, models.InstrumentRequest{InstrumentID: instrument})
	}

	reqBody := models.GetClosePricesRequest{Instruments: InstrumentRequests}
	url := fmt.Sprintf("%s/tinkoff.public.invest.api.contract.v1.MarketDataService/GetClosePrices", s.Config.APIBaseURL)

	headers := map[string]string{
		"Authorization": "Bearer " + s.Config.APIToken,
		"Content-Type":  "application/json",
	}

	respBody, err := s.Client.Post(url, headers, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.ClosePricesResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	start := time.Now()
	if err := s.KafkaProducer.SendMessage(s.KafkaProducer.Topic, respBody); err != nil {
		return nil, err
	}

	fmt.Println("Time taken for KafkaProducer.SendMessage:", time.Since(start))

	return response.ClosePrices, nil
}

func (s *TinkoffService) GetAllBonds(instrumentStatus string) ([]models.PlacementPrice, error) {
	reqBody := models.BondsRequest{InstrumentStatus: instrumentStatus}
	url := fmt.Sprintf("%s/tinkoff.public.invest.api.contract.v1.InstrumentsService/Shares", s.Config.APIBaseURL)

	headers := map[string]string{
		"Authorization": "Bearer " + s.Config.APIToken,
		"Content-Type":  "application/json",
	}

	respBody, err := s.Client.Post(url, headers, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetBondsResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response.Instruments, nil
}

func (s *TinkoffService) GetCandles(instrumentInfo map[string]any) ([]models.HistoricCandle, error) {
	reqBody := models.GetCandlesRequest{
		Figi: instrumentInfo["figi"].(string),
		From:  instrumentInfo["from"].(string),
		To: instrumentInfo["to"].(string),
		Interval: instrumentInfo["interval"].(string),
		InstrumentId: instrumentInfo["instrumentId"].(string),
	}

	url := fmt.Sprintf("%s/tinkoff.public.invest.api.contract.v1.MarketDataService/GetCandles", s.Config.APIBaseURL)

	headers := map[string]string{
		"Authorization": "Bearer " + s.Config.APIToken,
		"Content-Type":  "application/json",
	}

	respBody, err := s.Client.Post(url,headers,reqBody)
	if err != nil {
		return nil, err
	}

	var responce models.GetCandlesResponse
	if err := json.Unmarshal(respBody, &responce); err != nil{
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return responce.Candles, nil
}
