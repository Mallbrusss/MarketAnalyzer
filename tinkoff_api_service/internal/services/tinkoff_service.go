package services

import (
	"encoding/json"
	"fmt"
	"log"
	"tinkoff-api/config"
	"tinkoff-api/internal/kafka"
	"tinkoff-api/internal/models"
	"tinkoff-api/pkg/httpclient"
	"tinkoff-api/pkg/split"
)

const ChunkSize = 524288

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

	if err := s.KafkaProducer.SendMessage("closePriceData", "closePrice", respBody); err != nil {
		return nil, err
	}

	return response.ClosePrices, nil
}

func (s *TinkoffService) GetAllInstruments(instrumentStatus string) ([]models.PlacementPrice, error) {
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

	fmt.Printf("Response body size: %d bytes\n", len(respBody))
	chunks := split.SplitMessage(respBody, ChunkSize)
	if len(chunks) == 0 {
		log.Println("SplitMessage returned no chunks")
	}
	fmt.Printf("Total chunks: %d\n", len(chunks))

	var instrumentMap models.InstrumentPart

	for num, chunk := range chunks {
		key := fmt.Sprintf("allBonds:%d", num+1)

		instrumentMap.Data = chunk
		instrumentMap.MessageID = key
		instrumentMap.Part = num + 1
		instrumentMap.Total = len(chunks)

		data, err := json.Marshal(instrumentMap)
		if err != nil {
			return nil, err
		}

		err = s.KafkaProducer.SendMessage("bondsData", key, data)
		if err != nil {
			log.Printf("Failed to send chunk %d: %v", num+1, err)
			return nil, err
		}
		log.Printf("Sending chunk %d, size: %d bytes", num+1, len(chunk))
	}

	var response models.GetBondsResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response.Instruments, nil
}

func (s *TinkoffService) GetCandles(instrumentInfo map[string]any) ([]models.HistoricCandle, error) {
	reqBody := models.GetCandlesRequest{
		Figi:         instrumentInfo["figi"].(string),
		From:         instrumentInfo["from"].(string),
		To:           instrumentInfo["to"].(string),
		Interval:     instrumentInfo["interval"].(string),
		InstrumentId: instrumentInfo["instrumentId"].(string),
	}

	url := fmt.Sprintf("%s/tinkoff.public.invest.api.contract.v1.MarketDataService/GetCandles", s.Config.APIBaseURL)

	headers := map[string]string{
		"Authorization": "Bearer " + s.Config.APIToken,
		"Content-Type":  "application/json",
	}

	respBody, err := s.Client.Post(url, headers, reqBody)
	if err != nil {
		return nil, err
	}

	responce, data, err := s.fixeRespBody(respBody, reqBody.InstrumentId)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Response body size: %d bytes\n", len(respBody))
	chunks := split.SplitMessage(data, ChunkSize)
	if len(chunks) == 0 {
		log.Println("SplitMessage returned no chunks")
	}
	fmt.Printf("Total chunks: %d\n", len(chunks))

	var candlesPart models.InstrumentPart

	for num, chunk := range chunks {
		key := fmt.Sprintf("allBonds:%d", num+1)

		candlesPart.Data = chunk
		candlesPart.MessageID = key
		candlesPart.Part = num + 1
		candlesPart.Total = len(chunks)

		data, err := json.Marshal(candlesPart)
		if err != nil {
			return nil, err
		}
		err = s.KafkaProducer.SendMessage("candlesData", key, data)
		if err != nil {
			log.Printf("Failed to send chunk %d: %v", num+1, err)
			return nil, err
		}
		log.Printf("Sending chunk %d, size: %d bytes", num+1, len(chunk))
	}

	return responce.Candles, nil
}

func (s *TinkoffService) fixeRespBody(respBody []byte, instrumentID string) (models.GetCandlesResponse, []byte, error) {
	var responce models.GetCandlesResponse
	err := json.Unmarshal(respBody, &responce)
	if err != nil {
		return models.GetCandlesResponse{}, nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	for i := range responce.Candles {
		responce.Candles[i].InstrumentId = instrumentID
	}

	data, err := json.Marshal(responce)
	if err != nil {
		return models.GetCandlesResponse{}, nil, err
	}

	return responce, data, nil
}
