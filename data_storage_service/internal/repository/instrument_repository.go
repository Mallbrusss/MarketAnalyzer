package repository

import (
	"data-storage/internal/models"
	"log"

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

func (ir *InstrumentRepository) CreateInstruments(instruments []models.PlacementPrice) error {
	batchSize := 100

	for i := 0; i < len(instruments); i += batchSize {
		end := i + batchSize
		if end > len(instruments) {
			end = len(instruments)
		}

		batch := make([]*models.PlacementPrice, end-i)
		for j := 0; j < len(batch); j++ {
			batch[j] = &instruments[i+j]
		}

		tx := ir.db.Begin()
		if tx.Error != nil {
			log.Printf("failed to start transaction: %v", tx.Error)
			return tx.Error
		}

		batchNumber := (i / batchSize) + 1
		totalBatches := (len(instruments) + batchSize - 1) / batchSize

		log.Printf("Create batch %d of %d", batchNumber, totalBatches)

		err := tx.Create(&batch).Error
		if err != nil {
			tx.Rollback()
			log.Printf("failed to insert batch: %v", err)
			return err
		}

		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			log.Printf("failed to commit transaction: %v", err)
			return err
		}
	}

	log.Println("Instument create success")
	return nil
}
