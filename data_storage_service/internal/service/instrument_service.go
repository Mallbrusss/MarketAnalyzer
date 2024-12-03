package service

import (
	"data-storage/internal/models"
)

type InstrumentRepository interface {
	CreateInstruments(instruments []models.PlacementPrice) error
}

type InstrumentService struct {
	instrumentRepository InstrumentRepository
}

func NewInstrumentService(instrumentRepository InstrumentRepository) *InstrumentService {
	return &InstrumentService{
		instrumentRepository: instrumentRepository,
	}
}

func (is *InstrumentService) CreateInstruments(instruments []models.PlacementPrice) error {
	err := is.instrumentRepository.CreateInstruments(instruments)
	if err != nil {
		return err
	}
	return nil
}
