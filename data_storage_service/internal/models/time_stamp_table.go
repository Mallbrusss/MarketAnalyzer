package models

import "time"

type HistoricCandle struct {
	Volume string    `json:"volume"`
	High   High      `json:"high"`
	Low    Low       `json:"low"`
	Time   time.Time `json:"time"`
	Close  Close     `json:"close"`
	Open   Open      `json:"open"`
}

type High struct {
	Units string `json:"units"`
	Nano  int    `json:"nano"`
}

type Low struct {
	Units string `json:"units"`
	Nano  int    `json:"nano"`
}

type Close struct {
	Units string `json:"units"`
	Nano  int    `json:"nano"`
}

type Open struct {
	Units string `json:"units"`
	Nano  int    `json:"nano"`
}

type HistoricCandles struct {
	Candles []HistoricCandle `json:"candles"`
}

type Candle struct {
	InstrumentID string    `gorm:"primaryKey;size:255"`
	Timestamp    time.Time `gorm:"primaryKey"`
	Open         float64   `gorm:"type:numeric(18,4);not null"`
	High         float64   `gorm:"type:numeric(18,4);not null"`
	Low          float64   `gorm:"type:numeric(18,4);not null"`
	Close        float64   `gorm:"type:numeric(18,4);not null"`
	Volume       int64     `gorm:"not null"`
}
