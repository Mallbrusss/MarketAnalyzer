package service

import (
	"data-storage/internal/models"
	"data-storage/internal/repository"
)

type InstrumentRepository interface {
	CreateInstruments(instruments []models.Instrument) error
}

type InstrumentService struct {
	instrumentRepository InstrumentRepository
}

func NewInstrumentService(instrumentRepository *repository.InstrumentRepository) *InstrumentService {
	return &InstrumentService{
		instrumentRepository: instrumentRepository,
	}
}

func (is *InstrumentService) CreateInstruments(instruments []models.Instrument) error {
	err := is.instrumentRepository.CreateInstruments(instruments)
	if err != nil {
		return err
	}
	return nil
}
