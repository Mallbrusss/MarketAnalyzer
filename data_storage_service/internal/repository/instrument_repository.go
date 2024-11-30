package repository

import (
	"data-storage/internal/models"

	"gorm.io/gorm"
)

type InstrumentRepository struct {
	db *gorm.DB
}

func NewInstrumentRepository(db *gorm.DB) *InstrumentRepository {
	return &InstrumentRepository{
		db: db,
	}
}

func (ir *InstrumentRepository) CreateInstruments(instruments []models.Instrument) error {
	tx := ir.db.Begin()
	if tx.Error != nil{
		return tx.Error
	}

	err := tx.Create(instruments).Error
	if err != nil{
		return err
	}

	if err := tx.Commit().Error; err != nil{
		return err
	}

	return nil

}
